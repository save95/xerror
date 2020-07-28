package xcode

import "sync"

var _repo *repository
var once sync.Once

type repository struct {
	codes sync.Map
}

func Repository() *repository {
	once.Do(func() {
		_repo = &repository{}
		for _, code := range allCode {
			_repo.codes.Store(code.Code(), code)
		}
	})

	return _repo
}

func (r *repository) LoadCodes(codes ...XCode) {
	_repo.codes = sync.Map{}
	for _, code := range codes {
		r.codes.Store(code.Code(), code)
	}
}

func (r *repository) AppendCodes(codes ...XCode) {
	for _, code := range codes {
		if _, ok := r.codes.Load(code.Code()); !ok {
			r.codes.Store(code.Code(), code)
		}
	}
}
