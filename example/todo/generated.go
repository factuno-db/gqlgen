// This file was generated by github.com/vektah/gqlgen, DO NOT EDIT

package todo

import (
	context "context"
	strconv "strconv"
	sync "sync"

	graphql "github.com/vektah/gqlgen/graphql"
	errors "github.com/vektah/gqlgen/neelance/errors"
	introspection "github.com/vektah/gqlgen/neelance/introspection"
	query "github.com/vektah/gqlgen/neelance/query"
	schema "github.com/vektah/gqlgen/neelance/schema"
)

func MakeExecutableSchema(resolvers Resolvers) graphql.ExecutableSchema {
	return &executableSchema{resolvers}
}

type Resolvers interface {
	MyMutation_createTodo(ctx context.Context, todo TodoInput) (Todo, error)
	MyMutation_updateTodo(ctx context.Context, id int, changes map[string]interface{}) (*Todo, error)
	MyQuery_todo(ctx context.Context, id int) (*Todo, error)
	MyQuery_lastTodo(ctx context.Context) (*Todo, error)
	MyQuery_todos(ctx context.Context) ([]Todo, error)
}

type Todo struct {
	ID   int
	Text string
	Done bool
}

type TodoInput struct {
	Text string
	Done *bool
}

type executableSchema struct {
	resolvers Resolvers
}

func (e *executableSchema) Schema() *schema.Schema {
	return parsedSchema
}

func (e *executableSchema) Query(ctx context.Context, doc *query.Document, variables map[string]interface{}, op *query.Operation) *graphql.Response {
	ec := executionContext{resolvers: e.resolvers, variables: variables, doc: doc, ctx: ctx}

	data := ec._myQuery(op.Selections)
	ec.wg.Wait()

	return &graphql.Response{
		Data:   data,
		Errors: ec.Errors,
	}
}

func (e *executableSchema) Mutation(ctx context.Context, doc *query.Document, variables map[string]interface{}, op *query.Operation) *graphql.Response {
	ec := executionContext{resolvers: e.resolvers, variables: variables, doc: doc, ctx: ctx}

	data := ec._myMutation(op.Selections)
	ec.wg.Wait()

	return &graphql.Response{
		Data:   data,
		Errors: ec.Errors,
	}
}

func (e *executableSchema) Subscription(ctx context.Context, doc *query.Document, variables map[string]interface{}, op *query.Operation) <-chan *graphql.Response {
	events := make(chan *graphql.Response, 1)
	events <- &graphql.Response{Errors: []*errors.QueryError{{Message: "subscriptions are not supported"}}}
	return events
}

type executionContext struct {
	errors.Builder
	resolvers Resolvers
	variables map[string]interface{}
	doc       *query.Document
	ctx       context.Context
	wg        sync.WaitGroup
}

var myMutationImplementors = []string{"MyMutation"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _myMutation(sel []query.Selection) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, myMutationImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias
		out.Values[i] = graphql.Null

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("MyMutation")
		case "createTodo":
			var arg0 TodoInput
			if tmp, ok := field.Args["todo"]; ok {
				var err error

				arg0, err = UnmarshalTodoInput(tmp)
				if err != nil {
					ec.Error(err)
					return graphql.Null
				}
			}
			res, err := ec.resolvers.MyMutation_createTodo(ec.ctx, arg0)
			if err != nil {
				ec.Error(err)
				continue
			}

			out.Values[i] = ec._todo(field.Selections, &res)
		case "updateTodo":
			var arg0 int
			if tmp, ok := field.Args["id"]; ok {
				var err error

				arg0, err = graphql.UnmarshalInt(tmp)
				if err != nil {
					ec.Error(err)
					return graphql.Null
				}
			}
			var arg1 map[string]interface{}
			if tmp, ok := field.Args["changes"]; ok {
				var err error

				arg1, err = graphql.UnmarshalMap(tmp)
				if err != nil {
					ec.Error(err)
					return graphql.Null
				}
			}
			res, err := ec.resolvers.MyMutation_updateTodo(ec.ctx, arg0, arg1)
			if err != nil {
				ec.Error(err)
				continue
			}

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = ec._todo(field.Selections, res)
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

var myQueryImplementors = []string{"MyQuery"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _myQuery(sel []query.Selection) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, myQueryImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias
		out.Values[i] = graphql.Null

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("MyQuery")
		case "todo":
			var arg0 int
			if tmp, ok := field.Args["id"]; ok {
				var err error

				arg0, err = graphql.UnmarshalInt(tmp)
				if err != nil {
					ec.Error(err)
					return graphql.Null
				}
			}
			ec.wg.Add(1)
			go func(i int, field graphql.CollectedField) {
				defer ec.wg.Done()
				res, err := ec.resolvers.MyQuery_todo(ec.ctx, arg0)
				if err != nil {
					ec.Error(err)
					return
				}

				if res == nil {
					out.Values[i] = graphql.Null
				} else {
					out.Values[i] = ec._todo(field.Selections, res)
				}
			}(i, field)
		case "lastTodo":
			ec.wg.Add(1)
			go func(i int, field graphql.CollectedField) {
				defer ec.wg.Done()
				res, err := ec.resolvers.MyQuery_lastTodo(ec.ctx)
				if err != nil {
					ec.Error(err)
					return
				}

				if res == nil {
					out.Values[i] = graphql.Null
				} else {
					out.Values[i] = ec._todo(field.Selections, res)
				}
			}(i, field)
		case "todos":
			ec.wg.Add(1)
			go func(i int, field graphql.CollectedField) {
				defer ec.wg.Done()
				res, err := ec.resolvers.MyQuery_todos(ec.ctx)
				if err != nil {
					ec.Error(err)
					return
				}

				arr1 := graphql.Array{}
				for idx1 := range res {
					var tmp1 graphql.Marshaler
					tmp1 = ec._todo(field.Selections, &res[idx1])
					arr1 = append(arr1, tmp1)
				}
				out.Values[i] = arr1
			}(i, field)
		case "__schema":
			res := ec.introspectSchema()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = ec.___Schema(field.Selections, res)
			}
		case "__type":
			var arg0 string
			if tmp, ok := field.Args["name"]; ok {
				var err error

				arg0, err = graphql.UnmarshalString(tmp)
				if err != nil {
					ec.Error(err)
					return graphql.Null
				}
			}
			res := ec.introspectType(arg0)

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = ec.___Type(field.Selections, res)
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

var todoImplementors = []string{"Todo"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _todo(sel []query.Selection, it *Todo) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, todoImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias
		out.Values[i] = graphql.Null

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Todo")
		case "id":
			res := it.ID

			out.Values[i] = graphql.MarshalInt(res)
		case "text":
			res := it.Text

			out.Values[i] = graphql.MarshalString(res)
		case "done":
			res := it.Done

			out.Values[i] = graphql.MarshalBoolean(res)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

var __DirectiveImplementors = []string{"__Directive"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Directive(sel []query.Selection, it *introspection.Directive) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, __DirectiveImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias
		out.Values[i] = graphql.Null

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Directive")
		case "name":
			res := it.Name()

			out.Values[i] = graphql.MarshalString(res)
		case "description":
			res := it.Description()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = graphql.MarshalString(*res)
			}
		case "locations":
			res := it.Locations()

			arr1 := graphql.Array{}
			for idx1 := range res {
				var tmp1 graphql.Marshaler
				tmp1 = graphql.MarshalString(res[idx1])
				arr1 = append(arr1, tmp1)
			}
			out.Values[i] = arr1
		case "args":
			res := it.Args()

			arr1 := graphql.Array{}
			for idx1 := range res {
				var tmp1 graphql.Marshaler

				if res[idx1] == nil {
					tmp1 = graphql.Null
				} else {
					tmp1 = ec.___InputValue(field.Selections, res[idx1])
				}
				arr1 = append(arr1, tmp1)
			}
			out.Values[i] = arr1
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

var __EnumValueImplementors = []string{"__EnumValue"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___EnumValue(sel []query.Selection, it *introspection.EnumValue) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, __EnumValueImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias
		out.Values[i] = graphql.Null

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__EnumValue")
		case "name":
			res := it.Name()

			out.Values[i] = graphql.MarshalString(res)
		case "description":
			res := it.Description()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = graphql.MarshalString(*res)
			}
		case "isDeprecated":
			res := it.IsDeprecated()

			out.Values[i] = graphql.MarshalBoolean(res)
		case "deprecationReason":
			res := it.DeprecationReason()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = graphql.MarshalString(*res)
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

var __FieldImplementors = []string{"__Field"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Field(sel []query.Selection, it *introspection.Field) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, __FieldImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias
		out.Values[i] = graphql.Null

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Field")
		case "name":
			res := it.Name()

			out.Values[i] = graphql.MarshalString(res)
		case "description":
			res := it.Description()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = graphql.MarshalString(*res)
			}
		case "args":
			res := it.Args()

			arr1 := graphql.Array{}
			for idx1 := range res {
				var tmp1 graphql.Marshaler

				if res[idx1] == nil {
					tmp1 = graphql.Null
				} else {
					tmp1 = ec.___InputValue(field.Selections, res[idx1])
				}
				arr1 = append(arr1, tmp1)
			}
			out.Values[i] = arr1
		case "type":
			res := it.Type()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = ec.___Type(field.Selections, res)
			}
		case "isDeprecated":
			res := it.IsDeprecated()

			out.Values[i] = graphql.MarshalBoolean(res)
		case "deprecationReason":
			res := it.DeprecationReason()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = graphql.MarshalString(*res)
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

var __InputValueImplementors = []string{"__InputValue"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___InputValue(sel []query.Selection, it *introspection.InputValue) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, __InputValueImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias
		out.Values[i] = graphql.Null

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__InputValue")
		case "name":
			res := it.Name()

			out.Values[i] = graphql.MarshalString(res)
		case "description":
			res := it.Description()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = graphql.MarshalString(*res)
			}
		case "type":
			res := it.Type()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = ec.___Type(field.Selections, res)
			}
		case "defaultValue":
			res := it.DefaultValue()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = graphql.MarshalString(*res)
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

var __SchemaImplementors = []string{"__Schema"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Schema(sel []query.Selection, it *introspection.Schema) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, __SchemaImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias
		out.Values[i] = graphql.Null

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Schema")
		case "types":
			res := it.Types()

			arr1 := graphql.Array{}
			for idx1 := range res {
				var tmp1 graphql.Marshaler

				if res[idx1] == nil {
					tmp1 = graphql.Null
				} else {
					tmp1 = ec.___Type(field.Selections, res[idx1])
				}
				arr1 = append(arr1, tmp1)
			}
			out.Values[i] = arr1
		case "queryType":
			res := it.QueryType()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = ec.___Type(field.Selections, res)
			}
		case "mutationType":
			res := it.MutationType()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = ec.___Type(field.Selections, res)
			}
		case "subscriptionType":
			res := it.SubscriptionType()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = ec.___Type(field.Selections, res)
			}
		case "directives":
			res := it.Directives()

			arr1 := graphql.Array{}
			for idx1 := range res {
				var tmp1 graphql.Marshaler

				if res[idx1] == nil {
					tmp1 = graphql.Null
				} else {
					tmp1 = ec.___Directive(field.Selections, res[idx1])
				}
				arr1 = append(arr1, tmp1)
			}
			out.Values[i] = arr1
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

var __TypeImplementors = []string{"__Type"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Type(sel []query.Selection, it *introspection.Type) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, __TypeImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias
		out.Values[i] = graphql.Null

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Type")
		case "kind":
			res := it.Kind()

			out.Values[i] = graphql.MarshalString(res)
		case "name":
			res := it.Name()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = graphql.MarshalString(*res)
			}
		case "description":
			res := it.Description()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = graphql.MarshalString(*res)
			}
		case "fields":
			var arg0 bool
			if tmp, ok := field.Args["includeDeprecated"]; ok {
				var err error

				arg0, err = graphql.UnmarshalBoolean(tmp)
				if err != nil {
					ec.Error(err)
					return graphql.Null
				}
			}
			res := it.Fields(arg0)

			arr1 := graphql.Array{}
			for idx1 := range res {
				var tmp1 graphql.Marshaler

				if res[idx1] == nil {
					tmp1 = graphql.Null
				} else {
					tmp1 = ec.___Field(field.Selections, res[idx1])
				}
				arr1 = append(arr1, tmp1)
			}
			out.Values[i] = arr1
		case "interfaces":
			res := it.Interfaces()

			arr1 := graphql.Array{}
			for idx1 := range res {
				var tmp1 graphql.Marshaler

				if res[idx1] == nil {
					tmp1 = graphql.Null
				} else {
					tmp1 = ec.___Type(field.Selections, res[idx1])
				}
				arr1 = append(arr1, tmp1)
			}
			out.Values[i] = arr1
		case "possibleTypes":
			res := it.PossibleTypes()

			arr1 := graphql.Array{}
			for idx1 := range res {
				var tmp1 graphql.Marshaler

				if res[idx1] == nil {
					tmp1 = graphql.Null
				} else {
					tmp1 = ec.___Type(field.Selections, res[idx1])
				}
				arr1 = append(arr1, tmp1)
			}
			out.Values[i] = arr1
		case "enumValues":
			var arg0 bool
			if tmp, ok := field.Args["includeDeprecated"]; ok {
				var err error

				arg0, err = graphql.UnmarshalBoolean(tmp)
				if err != nil {
					ec.Error(err)
					return graphql.Null
				}
			}
			res := it.EnumValues(arg0)

			arr1 := graphql.Array{}
			for idx1 := range res {
				var tmp1 graphql.Marshaler

				if res[idx1] == nil {
					tmp1 = graphql.Null
				} else {
					tmp1 = ec.___EnumValue(field.Selections, res[idx1])
				}
				arr1 = append(arr1, tmp1)
			}
			out.Values[i] = arr1
		case "inputFields":
			res := it.InputFields()

			arr1 := graphql.Array{}
			for idx1 := range res {
				var tmp1 graphql.Marshaler

				if res[idx1] == nil {
					tmp1 = graphql.Null
				} else {
					tmp1 = ec.___InputValue(field.Selections, res[idx1])
				}
				arr1 = append(arr1, tmp1)
			}
			out.Values[i] = arr1
		case "ofType":
			res := it.OfType()

			if res == nil {
				out.Values[i] = graphql.Null
			} else {
				out.Values[i] = ec.___Type(field.Selections, res)
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func UnmarshalTodoInput(v interface{}) (TodoInput, error) {
	var it TodoInput

	for k, v := range v.(map[string]interface{}) {
		switch k {
		case "text":
			var err error

			it.Text, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "done":
			var err error
			var ptr1 bool

			ptr1, err = graphql.UnmarshalBoolean(v)
			it.Done = &ptr1
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

var parsedSchema = schema.MustParse("schema {\n\tquery: MyQuery\n\tmutation: MyMutation\n}\n\ntype MyQuery {\n\ttodo(id: Int!): Todo\n\tlastTodo: Todo\n\ttodos: [Todo!]!\n}\n\ntype MyMutation {\n\tcreateTodo(todo: TodoInput!): Todo!\n\tupdateTodo(id: Int!, changes: Map!): Todo\n}\n\ntype Todo {\n\tid: Int!\n\ttext: String!\n\tdone: Boolean!\n}\n\ninput TodoInput {\n\ttext: String!\n\tdone: Boolean\n}\n")

func (ec *executionContext) introspectSchema() *introspection.Schema {
	return introspection.WrapSchema(parsedSchema)
}

func (ec *executionContext) introspectType(name string) *introspection.Type {
	t := parsedSchema.Resolve(name)
	if t == nil {
		return nil
	}
	return introspection.WrapType(t)
}