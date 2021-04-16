package crawl

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/nghiant3223/tikihackathon/internal/model"
	"github.com/nghiant3223/tikihackathon/pkg/log"
	"github.com/spf13/cast"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
)

type Crawler struct {
	db         *gorm.DB
	cfg        *config
	httpClient *http.Client
}

func NewCrawler(db *gorm.DB, httpClient *http.Client, configFns ...configFn) *Crawler {
	cfg := &config{}
	for _, fn := range configFns {
		fn(cfg)
	}
	return &Crawler{
		db:         db,
		cfg:        cfg,
		httpClient: httpClient,
	}
}

func (c *Crawler) Start(ctx context.Context) error {
	err := c.crawlAll(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (c *Crawler) crawlAll(ctx context.Context) error {
	var errGroup errgroup.Group
	pool := make(chan struct{}, c.cfg.concurrency)
	for i := 0; i < c.cfg.count; i++ {
		recipeID := rand.Int() % c.cfg.upperID
		errGroup.Go(func() error {
			pool <- struct{}{}
			defer func() { <-pool }()
			resp, err := c.crawlRecipe(ctx, recipeID)
			if err != nil {
				log.Errorw("cannot craw single recipe", "recipe_id", recipeID, "error", err)
				return err
			}
			log.Infow("data", "resp", resp)
			err = c.persistData(ctx, resp)
			if err != nil && !isErrDuplicatedRecord(err) {
				log.Errorw("cannot persist crawl data ", "crawl_data", resp, "error", err)
				return err
			}
			return nil
		})
	}
	return errGroup.Wait()
}

func (c *Crawler) persistData(ctx context.Context, data Data) error {
	recipe := &model.Recipe{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		Rating:      data.Rating,
		Servings:    data.Servings,
		TotalTime:   data.TotalTime,
		Difficulty:  rand.Int()%10 + 1,
	}
	err := c.upsertRecipe(ctx, recipe)
	if err != nil {
		log.Errorw("cannot upsert recipe", "error", err)
		return err
	}

	for _, crawledIngredient := range data.Ingredients {
		if crawledIngredient.Name == nil {
			continue
		}

		ingredient := &model.Ingredient{
			Name: *crawledIngredient.Name,
		}
		err = c.upsertIngredient(ctx, ingredient)
		if err != nil {
			log.Errorw("cannot upsert ingredient", "error", err)
			continue
		}

		unit := &model.Unit{
			Name: crawledIngredient.Unit.Unit,
			ID:   cast.ToInt(crawledIngredient.Unit.Value),
		}
		err = c.upsertUnit(ctx, unit)
		if err != nil {
			log.Errorw("cannot upsert unit", "error", err)
			continue
		}

		including := &model.Including{
			RecipeID:     recipe.ID,
			IngredientID: ingredient.ID,
			UnitID:       unit.ID,
			Quantity:     crawledIngredient.Quantity,
		}
		err = c.upsertIncluding(ctx, including)
		if err != nil {
			log.Errorw("cannot upsert including", "error", err)
			continue
		}

	}

	for _, crawledStep := range data.Steps {
		step := &model.Step{
			Content:  crawledStep.Content,
			RecipeID: recipe.ID,
		}
		err = c.upsertStep(ctx, step)
		if err != nil {
			log.Errorw("cannot upsert step", "error", err)
			continue
		}

		for _, crawledPhoto := range crawledStep.Photos {
			largestCrawledPhoto := crawledPhoto[len(crawledPhoto)-1]
			photo := &model.StepPhoto{
				URL:    largestCrawledPhoto.URL,
				Height: largestCrawledPhoto.Height,
				Width:  largestCrawledPhoto.Width,
				StepID: step.ID,
			}
			err = c.upsertStepPhoto(ctx, photo)
			if err != nil {
				log.Errorw("cannot upsert step photo", "error", err)
				continue
			}
		}
	}

	for _, crawledPhoto := range data.Photos {
		largestCrawledPhoto := crawledPhoto[len(crawledPhoto)-1]
		photo := &model.RecipePhoto{
			URL:      largestCrawledPhoto.URL,
			Height:   largestCrawledPhoto.Height,
			Width:    largestCrawledPhoto.Width,
			RecipeID: recipe.ID,
		}
		err = c.upsertRecipePhoto(ctx, photo)
		if err != nil {
			log.Errorw("cannot upsert recipe photo", "error", err)
			continue
		}
	}

	return nil
}

func (c *Crawler) crawlRecipe(ctx context.Context, recipeID int) (Data, error) {
	recipeURL, err := c.buildTargetURL(recipeID)
	if err != nil {
		return Data{}, err
	}
	resp, err := c.makeCrawlRequest(ctx, recipeURL)
	if err != nil {
		return Data{}, err
	}
	return resp, nil
}

func (c *Crawler) buildTargetURL(id int) (string, error) {
	return c.cfg.target + "&id=" + cast.ToString(id), nil
	u, err := url.Parse(c.cfg.target)
	if err != nil {
		return "", err
	}
	query := u.Query()
	log.Info(cast.ToString(id))
	query.Add("id", cast.ToString(id))
	u.RawQuery = query.Encode()
	return u.String(), nil
}

func (c *Crawler) makeCrawlRequest(ctx context.Context, url string) (Data, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return Data{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil || resp.StatusCode >= 300 {
		return Data{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Data{}, err
	}
	var crawlResponse Response
	err = json.Unmarshal(body, &crawlResponse)
	if err != nil {
		return Data{}, err
	}
	if crawlResponse.Code < 0 {
		return Data{}, errors.New("error occurs when calling api")
	}
	return crawlResponse.Data, nil
}

func (c *Crawler) upsertRecipe(ctx context.Context, recipe *model.Recipe) error {
	return c.db.
		WithContext(ctx).
		Create(recipe).Error
}

func (c *Crawler) upsertIngredient(ctx context.Context, ingredient *model.Ingredient) error {
	err := c.db.
		WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "name"}},
			DoNothing: true,
		}).
		Create(ingredient).Error
	if err != nil {
		return err
	}
	if ingredient.ID == 0 {
		return errDuplicatedRecord
	}
	return nil
}

func (c *Crawler) upsertUnit(ctx context.Context, unit *model.Unit) error {
	err := c.db.
		WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "name"}},
			DoNothing: true,
		}).
		Create(unit).Error
	if err != nil {
		return err
	}
	if unit.ID == 0 {
		return errDuplicatedRecord
	}
	return nil
}

func (c *Crawler) upsertIncluding(ctx context.Context, including *model.Including) error {
	return c.db.
		WithContext(ctx).
		Create(including).Error
}

func (c *Crawler) upsertStep(ctx context.Context, step *model.Step) error {
	return c.db.
		WithContext(ctx).
		Create(step).Error
}

func (c *Crawler) upsertStepPhoto(ctx context.Context, photo *model.StepPhoto) error {
	return c.db.
		WithContext(ctx).
		Create(photo).Error
}

func (c *Crawler) upsertRecipePhoto(ctx context.Context, photo *model.RecipePhoto) error {
	return c.db.
		WithContext(ctx).
		Create(photo).Error
}

func (c *Crawler) Stop(ctx context.Context) error {
	return nil
}
