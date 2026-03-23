# Go Backend Development Learning Roadmap

## 🎯 Goal
Build a simple REST API application in Go from scratch, understanding the fundamentals and best practices.

---

## 📚 Phase 1: Go Fundamentals (Week 1)

### 1.1 Getting Started
- [x] Install Go and set up development environment
- [x] Understand Go workspace structure (GOPATH vs Go modules)
- [x] Write your first "Hello, World!" program
- [x] Understand `go run`, `go build`, `go install` commands
- [x] Set up Go modules (`go mod init`)

**Key Concepts:**
- Go is compiled (unlike Python/JS which are interpreted)
- Statically typed (unlike Python/JS which are dynamically typed)
- Fast compilation and execution
- Strong concurrency support built-in

### 1.2 Basic Syntax & Types
- [x] Variables and constants (`var`, `const`, `:=` shorthand)
- [x] Basic data types (int, float64, string, bool)
- [x] Type inference and explicit typing
- [x] Zero values (what happens when you declare without initializing)
- [x] Type conversions (explicit, no implicit conversions like JS)

**Differences from JS/Python:**
- No `null` - Go has `nil` for pointers, slices, maps, channels, functions, interfaces
- No truthy/falsy - must use explicit boolean comparisons
- No automatic type coercion

### 1.3 Control Flow
- [x] If/else statements (no parentheses needed!)
- [x] Switch statements (more powerful than JS/Python)
- [x] For loops (only loop type in Go - no while/do-while)
- [x] Range loops (for iterating over slices, maps, strings, channels)

### 1.4 Functions
- [x] Function declaration and calling
- [x] Multiple return values (common in Go!)
- [x] Named return values
- [x] Variadic functions (`...` syntax)
- [x] Functions as first-class citizens
- [x] Anonymous functions and closures

**Key Difference:**
- Go functions can return multiple values (like tuples in Python, but built-in)
- Error handling is typically done via return values, not exceptions

### 1.5 Packages and Imports
- [x] Understanding packages (`package` keyword)
- [x] Importing packages (`import`)
- [x] Exported vs unexported (capitalized vs lowercase)
- [x] Standard library packages (fmt, os, strings, etc.)
- [x] Creating your own packages

---

## 📚 Phase 2: Go-Specific Features (Week 1-2)

### 2.1 Pointers
- [x] Understanding pointers (`*` and `&`)
- [x] When to use pointers vs values
- [x] Nil pointers and how to check them
- [x] Pointer receivers in methods

**Important:**
- Go has pointers (like C), but no pointer arithmetic
- Everything is passed by value by default (unlike Python where objects are passed by reference)

### 2.2 Structs
- [x] Defining structs
- [x] Creating and initializing structs
- [x] Struct fields and methods
- [x] Pointer receivers vs value receivers
- [x] Struct embedding (composition over inheritance)

**Key Concept:**
- Go doesn't have classes - structs + methods = similar functionality
- No inheritance, but composition via embedding

### 2.3 Arrays and Slices
- [x] Arrays (fixed size, rarely used)
- [x] Slices (dynamic arrays, what you'll use most)
- [x] Slice operations (append, copy, slicing)
- [x] Slice internals (length, capacity, underlying array)
- [x] Making slices with `make()`

**Important:**
- Slices are references to underlying arrays
- Appending can cause reallocation
- Be careful with slice capacity

### 2.4 Maps
- [x] Creating and initializing maps
- [x] Accessing, adding, deleting map elements
- [x] Checking if key exists
- [x] Iterating over maps
- [x] Maps are reference types

**Difference from Python dicts:**
- Keys must be comparable types
- No `None` as default - use `ok` idiom to check existence

### 2.5 Interfaces
- [x] Understanding interfaces (implicit implementation)
- [x] Interface types and values
- [x] Empty interface (`interface{}` or `any`)
- [x] Type assertions and type switches
- [x] Interface composition

**Key Concept:**
- "If it quacks like a duck, it's a duck" - no explicit `implements` keyword
- Interfaces are satisfied implicitly

---

## 📚 Phase 3: Error Handling & Concurrency (Week 2)

### 3.1 Error Handling
- [x] Error type and creating errors
- [x] Error handling patterns (check, handle, return)
- [x] Custom error types
- [x] Error wrapping (`fmt.Errorf` with `%w`)
- [x] `errors.Is()` and `errors.As()`

**Key Difference:**
- No try/catch - errors are values
- Always check errors explicitly
- "If you ignore an error, you're asking for trouble"

### 3.2 Goroutines
- [x] Understanding goroutines (`go` keyword)
- [x] Creating concurrent functions
- [x] Goroutine lifecycle
- [x] When goroutines finish (or don't)

**Key Concept:**
- Goroutines are lightweight threads
- Much cheaper than OS threads
- Can spawn thousands easily

### 3.3 Channels
- [ ] Creating channels (`make(chan type)`)
- [ ] Sending and receiving (`<-` operator)
- [ ] Buffered vs unbuffered channels
- [ ] Channel direction (send-only, receive-only)
- [ ] Closing channels
- [ ] Range over channels
- [ ] Select statement (like switch for channels)

**Key Concept:**
- Channels are Go's way of communicating between goroutines
- "Don't communicate by sharing memory; share memory by communicating"

### 3.4 Sync Package
- [ ] WaitGroups (waiting for goroutines)
- [ ] Mutexes (mutual exclusion locks)
- [ ] RWMutex (read-write locks)
- [ ] Once (executing code once)
- [ ] Atomic operations

---

## 📚 Phase 4: Working with Data (Week 2-3)

### 4.1 JSON Handling
- [ ] Encoding/decoding JSON (`encoding/json`)
- [ ] Struct tags for JSON (`json:"field_name"`)
- [ ] Marshaling and unmarshaling
- [ ] Handling nested structures
- [ ] Custom JSON marshaling/unmarshaling

### 4.2 File I/O
- [ ] Reading files (`os.ReadFile`, `ioutil` - deprecated, use `os`)
- [ ] Writing files
- [ ] Working with file paths (`path/filepath`)
- [ ] Reading line by line

### 4.3 Working with Strings
- [ ] String package functions
- [ ] String builders (`strings.Builder`)
- [ ] String formatting (`fmt.Sprintf`)
- [ ] Regular expressions (`regexp`)

---

## 📚 Phase 5: Building Your First API (Week 3)

### 5.1 HTTP Basics in Go
- [ ] Understanding `net/http` package
- [ ] HTTP handlers and handler functions
- [ ] HTTP methods (GET, POST, PUT, DELETE)
- [ ] Request and Response objects
- [ ] Reading request body
- [ ] Writing response body
- [ ] Status codes

### 5.2 Routing
- [ ] Using `http.ServeMux` (built-in router)
- [ ] URL patterns and path matching
- [ ] Query parameters
- [ ] Path parameters (manual parsing or router library)

**Note:** For production, you'll likely use a router like `gorilla/mux` or `chi`, but start with standard library!

### 5.3 Middleware
- [ ] Understanding middleware concept
- [ ] Creating middleware functions
- [ ] Chaining middleware
- [ ] Common middleware (logging, CORS, authentication)

### 5.4 Project Structure
- [ ] Organizing Go projects
- [ ] Separating handlers, models, services
- [ ] Package organization
- [ ] Best practices for project layout

---

## 📚 Phase 6: Advanced API Features (Week 3-4)

### 6.1 Configuration
- [ ] Environment variables (`os.Getenv`)
- [ ] Configuration files
- [ ] Using `flag` package for CLI flags
- [ ] Struct-based configuration

### 6.2 Database Integration
- [ ] Understanding `database/sql` package
- [ ] Connecting to databases (PostgreSQL, MySQL, SQLite)
- [ ] Executing queries
- [ ] Prepared statements
- [ ] Transactions
- [ ] Connection pooling

**Popular Libraries:**
- `sqlx` - extends database/sql
- `gorm` - ORM (like Django ORM, Sequelize)
- `pgx` - PostgreSQL driver

### 6.3 Validation
- [ ] Input validation
- [ ] Struct validation
- [ ] Using validation libraries (`validator`)

### 6.4 Authentication & Authorization
- [ ] JWT tokens
- [ ] Password hashing (`bcrypt`, `golang.org/x/crypto`)
- [ ] Session management
- [ ] Middleware for auth

---

## 📚 Phase 7: Testing & Best Practices (Week 4)

### 7.1 Testing
- [ ] Writing tests (`*_test.go` files)
- [ ] `testing` package
- [ ] Table-driven tests
- [ ] Testing HTTP handlers
- [ ] Mocking
- [ ] Benchmarking

### 7.2 Error Handling Best Practices
- [ ] Proper error propagation
- [ ] Error context
- [ ] Logging errors
- [ ] Error types for different scenarios

### 7.3 Code Organization
- [ ] Project structure patterns
- [ ] Dependency injection
- [ ] Separation of concerns
- [ ] Clean architecture basics

### 7.4 Documentation
- [ ] Writing Go doc comments
- [ ] Generating documentation (`go doc`)
- [ ] README best practices

---

## 📚 Phase 8: Production Readiness (Week 4-5)

### 8.1 Logging
- [ ] Using `log` package
- [ ] Structured logging (`logrus`, `zap`)
- [ ] Log levels
- [ ] Logging best practices

### 8.2 Graceful Shutdown
- [ ] Handling OS signals
- [ ] Graceful server shutdown
- [ ] Cleanup resources

### 8.3 Performance
- [ ] Profiling Go applications
- [ ] Understanding allocations
- [ ] Optimizing hot paths
- [ ] Using `pprof`

### 8.4 Deployment
- [ ] Building for production
- [ ] Docker basics for Go apps
- [ ] Environment configuration
- [ ] Health checks

---

## 🎯 Final Project: Simple API App

### Project Requirements
Build a REST API with the following features:

1. **User Management**
   - Create user (POST /users)
   - Get user (GET /users/:id)
   - List users (GET /users)
   - Update user (PUT /users/:id)
   - Delete user (DELETE /users/:id)

2. **Data Storage**
   - Use in-memory storage (map) initially
   - Later: Add database (SQLite or PostgreSQL)

3. **Features**
   - JSON request/response
   - Input validation
   - Error handling
   - Basic logging
   - Health check endpoint

4. **Optional Enhancements**
   - Authentication
   - Pagination
   - Filtering/searching
   - Rate limiting

---

## 📖 Learning Resources

### Official
- [Go Tour](https://go.dev/tour/) - Interactive tutorial
- [Effective Go](https://go.dev/doc/effective_go) - Best practices
- [Go by Example](https://gobyexample.com/) - Code examples

### Books
- "The Go Programming Language" by Donovan & Kernighan
- "Go in Action" by William Kennedy

### Practice
- [Exercism Go Track](https://exercism.org/tracks/go)
- [LeetCode with Go](https://leetcode.com/)

---

## 💡 Tips for JS/Python Developers

1. **Type System**: Go is statically typed. You'll need to be explicit about types, but the compiler catches errors early.

2. **No Exceptions**: Errors are values. Always check them. It might feel verbose at first, but it makes error handling explicit.

3. **Concurrency**: Goroutines and channels are Go's superpower. They're easier than async/await in JS or threading in Python.

4. **Simplicity**: Go favors simplicity over cleverness. If something feels verbose, it's probably intentional.

5. **Compilation**: You'll catch many errors at compile time, not runtime. This is a feature, not a bug!

6. **Package Management**: Go modules (similar to npm/pip) handle dependencies. Use `go get` to add packages.

7. **Naming**: Exported (public) names start with capital letters. Unexported (private) start with lowercase.

---

## ✅ Progress Tracker

### Week 1
- [ ] Completed Phase 1
- [ ] Completed Phase 2
- [ ] Built small practice programs

### Week 2
- [ ] Completed Phase 3
- [ ] Completed Phase 4
- [ ] Built concurrent programs

### Week 3
- [ ] Completed Phase 5
- [ ] Completed Phase 6
- [ ] Started API project

### Week 4
- [ ] Completed Phase 7
- [ ] Completed Phase 8
- [ ] Finished API project

---

## 🚀 Next Steps After This Roadmap

Once you've completed this roadmap, you can explore:
- Microservices architecture
- gRPC
- WebSockets
- Advanced database patterns
- Kubernetes deployment
- Service mesh
- Performance optimization

---

**Remember:** Learning is a journey. Take your time with each concept. Write code, break things, fix them. That's how you learn! 🎓

Good luck on your Go journey! 🚀

