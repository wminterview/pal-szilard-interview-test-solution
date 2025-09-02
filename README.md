### Run with:
```powershell
docker compose up --build
```

### API:

#### **Books Management**
- `GET /api/books` - List all books
- `GET /api/books/:id` - Get details of a book
- `POST /api/books` - Create a new book
- `PUT /api/books/:id` - Update book details
- `DELETE /api/books/:id` - Delete a book

#### **Borrowing Management**
- `POST /api/books/:id/borrow` - Borrow a book
- `POST /api/books/:id/return` - Return a book
- `GET /api/borrowings` - List active borrowings

#### **Search & Filter**
- `GET /api/books?search=title` - Search books by title
- `GET /api/books?available=true` - List only available books