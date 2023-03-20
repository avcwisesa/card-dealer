package domain

type Store struct {
	dataStore map[string]Deck
}

func (s *Store) Get(id string) Deck {
	deck := s.dataStore[id]
	return deck
}

func (s *Store) AddOrUpdate(deck Deck) Deck {
	s.dataStore[deck.GetID()] = deck
	return deck
}

func NewStore() *Store {
	return &Store{
		dataStore: map[string]Deck{},
	}
}
