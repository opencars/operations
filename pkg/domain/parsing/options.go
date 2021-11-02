package parsing

func (mr *MapReduce) WithMapper(mapper Mapper) *MapReduce {
	mr.mapper = mapper
	return mr
}

func (mr *MapReduce) WithReducer(reducer Reducer) *MapReduce {
	mr.reducer = reducer
	return mr
}

func (mr *MapReduce) WithShuffler(shuffler Shuffler) *MapReduce {
	mr.shuffler = shuffler
	return mr
}
