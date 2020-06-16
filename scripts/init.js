// This script is used to populate the database with some initial data
db.auth('testUser', 'testPassword');
db = db.getSiblingDB('books');

db.books.drop();
db.books.save({
    title: 'The Great Gatsby',
    author: 'F. Scott Fitzgerald'
});
db.books.save({
    title: 'Moby Dick',
    author: 'Herman Melville'
});
db.books.save({
    title: 'War and Peace',
    author: 'Leo Tolstoy'
});