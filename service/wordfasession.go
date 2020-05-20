// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package service

import (
	"CiFa/wordfa"
	"sync"
	"time"
)

type WordfaSession struct {
	Task          *wordfa.Task
	SortAlgorithm int
	createAt      time.Time
	Resetting     bool
}

func NewWordfaSession(task *wordfa.Task, sortAlgorithm int) *WordfaSession {
	return &WordfaSession{
		Task:          task,
		SortAlgorithm: sortAlgorithm,
		createAt:      time.Now(),
	}
}

type WordFaSessionHolder struct {
	sessionMap map[string]*WordfaSession
	mux        sync.Mutex
}

func NewWordFaSessionHolder() *WordFaSessionHolder {
	return &WordFaSessionHolder{
		sessionMap: map[string]*WordfaSession{},
	}
}

func (w *WordFaSessionHolder) Put(token string, s *WordfaSession) {
	w.mux.Lock()
	defer w.mux.Unlock()

	w.sessionMap[token] = s
}

func (w *WordFaSessionHolder) Get(token string) (session *WordfaSession, ok bool) {
	w.mux.Lock()
	defer w.mux.Unlock()

	session, ok = w.sessionMap[token]
	return session, ok
}

func (w *WordFaSessionHolder) Reset(token string) {
	w.mux.Lock()
	defer w.mux.Unlock()

	if _, ok := w.sessionMap[token]; ok {
		w.sessionMap[token].Resetting = true
	}
}
