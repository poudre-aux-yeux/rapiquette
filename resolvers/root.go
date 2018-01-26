package resolvers

// NewRoot : TODO: link database
func NewRoot() (*RootResolver, error) {
	return &RootResolver{}, nil
}

// RootResolver : default resolver
type RootResolver struct{}
