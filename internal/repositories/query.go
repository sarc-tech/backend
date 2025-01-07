package repositories

import (
	"fmt"
	"strings"
)

type QueryOptions struct {
	AndFilters []string
	OrFilters  []string
	Args       []interface{}
	Limit      int
	Offset     int
}

type QueryOption func(*QueryOptions)

// Метод для обработки фильтров и построения WHERE-условий
func (q *QueryOptions) BuildConditions() (string, []interface{}) {
	conditions := []string{}
	argIndex := 1 // Начинаем нумерацию параметров с $1

	// Обрабатываем AND-фильтры
	if len(q.AndFilters) > 0 {
		for _, filter := range q.AndFilters {
			conditions = append(conditions, filter)
			argIndex++
		}
	}

	// Обрабатываем OR-фильтры
	if len(q.OrFilters) > 0 {
		for _, filter := range q.OrFilters {
			conditions = append(conditions, filter)
			argIndex++
		}
	}

	// Объединяем условия
	if len(conditions) > 0 {
		return " WHERE " + strings.Join(conditions, " AND "), q.Args
	}

	return "", q.Args
}

// Метод для добавления LIMIT и OFFSET
func (q *QueryOptions) AddLimitAndOffset(query string) string {
	if q.Limit > 0 {
		query += fmt.Sprintf(" LIMIT %d", q.Limit)
	}
	if q.Offset > 0 {
		query += fmt.Sprintf(" OFFSET %d", q.Offset)
	}
	return query
}

// Добавление фильтров и других опций
func WithAndFilter(column, operator string, value interface{}) QueryOption {
	return func(q *QueryOptions) {
		// Формируем строку фильтра
		filter := fmt.Sprintf("%s %s $%d", column, operator, len(q.Args)+1)
		q.AndFilters = append(q.AndFilters, filter)
		q.Args = append(q.Args, value)
	}
}

func WithOrFilter(column, operator string, value interface{}) QueryOption {
	return func(q *QueryOptions) {
		// Формируем строку фильтра
		filter := fmt.Sprintf("%s %s $%d", column, operator, len(q.Args)+1)
		q.OrFilters = append(q.OrFilters, filter)
		q.Args = append(q.Args, value)
	}
}

func WithLimit(limit int) QueryOption {
	return func(q *QueryOptions) {
		q.Limit = limit
	}
}

func WithOffset(offset int) QueryOption {
	return func(q *QueryOptions) {
		q.Offset = offset
	}
}
