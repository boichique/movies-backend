package movies

import (
	"context"

	"github.com/boichique/movie-reviews/internal/log"
	"github.com/boichique/movie-reviews/internal/modules/genres"
	"github.com/boichique/movie-reviews/internal/modules/stars"
	"golang.org/x/sync/errgroup"
)

type Service struct {
	repo          *Repository
	genresService *genres.Service
	starsService  *stars.Service
}

func NewService(repo *Repository, genresService *genres.Service, starsService *stars.Service) *Service {
	return &Service{
		repo:          repo,
		genresService: genresService,
		starsService:  starsService,
	}
}

func (s *Service) Create(ctx context.Context, movie *MovieDetails) error {
	if err := s.repo.Create(ctx, movie); err != nil {
		return err
	}

	log.FromContext(ctx).Info(
		"movie created",
		"movieID", movie.ID,
		"movieTitle", movie.Title)

	return s.assemble(ctx, movie)
}

func (s *Service) GetMoviesPaginated(ctx context.Context, searchTerm *string, starID *int, sortByRating *string, offset int, limit int) ([]*Movie, int, error) {
	return s.repo.GetMoviesPaginated(ctx, searchTerm, starID, sortByRating, offset, limit)
}

func (s *Service) GetByID(ctx context.Context, movieID int) (movie *MovieDetails, err error) {
	m, err := s.repo.GetByID(ctx, movieID)
	if err != nil {
		return nil, err
	}

	err = s.assemble(ctx, m)
	return m, err
}

func (s *Service) Update(ctx context.Context, movie *MovieDetails) error {
	if err := s.repo.Update(ctx, movie); err != nil {
		return err
	}

	log.FromContext(ctx).Info("movie updated",
		"movieTitle", movie.Title)
	return nil
}

func (s *Service) Delete(ctx context.Context, movieID int) error {
	if err := s.repo.Delete(ctx, movieID); err != nil {
		return err
	}

	log.FromContext(ctx).Info("movie deleted",
		"movieID", movieID)
	return nil
}

func (s *Service) assemble(ctx context.Context, movie *MovieDetails) error {
	group, groupCtx := errgroup.WithContext(ctx)

	group.Go(func() error {
		var err error
		movie.Genres, err = s.genresService.GetByMovieID(groupCtx, movie.ID)
		return err
	})
	group.Go(func() error {
		var err error
		movie.Cast, err = s.starsService.GetByMovieID(groupCtx, movie.ID)
		return err
	})

	return group.Wait()
}
