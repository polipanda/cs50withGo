## Week 9 Problems
https://cs50.harvard.edu/x/psets/9/

Complete the following exercises using the HTML, CSS, JavaScript, Python, and SQL.

### Birthdays

Complete a web application to keep track of friends' birthdays. Files are provided.
- In `app.py`, you'll find the start of a Flask web application with GET and POST routes to `/`
- `birthdays.db` is a SQLite database with one table for housing birthdays
- The `static` directory has a `styles.css` file for optional styling
- The `templates` directory has a `index.html` file for the frontend html/js

#### Specification
1. When the `/` route is requested via `GET`, a table should be rendered displaying all people in the database with 
their birthdays
    - In `app.py`, add logic so the `GET` queries the DB for birthdays and passes them to the `index.html` template
    - In `index.html`, add logic to render each birthday as a row in the table. The table should have two columns: one
   for the persons name, and one for their birthday
2. When the `/` route is requested, a new birthday should be added to the database, and the index page is re-rendered
    - In `index.html`, add a form for submitting a new birthday. It should let a user type in a name, a birthday month,
   and a birthday day
    - In `app.py`, add logic to handle a `POST` request that inserts a new row in the birthdays table of `birthday.db`
   using the data supplied by the form
3. Optionally
    - Add the ability to edit and/or delete birthday entries

### Finance
Complete a web application via which users `buy` and `sell` stocks. Most files are provided.
- A `finance.db` SQLite database which contains a users table. No users are created yet. New users cash amount will
default to 10,000.
- `app.py` contains stubs for every route that will be needed for this exercise
- `helpers.py` contains a `lookup` function for getting a stocks price, an `apology` function for rendering an error
page, and a `login_required` function for decorated routes to require a user be logged in to access the route
- `requirements.txt` lists the required packages for this exercise
- `static/` holds the styles.css file
- `templates/` will hold all the .html files for this exercise
    - note, in most cases, you will have to create a new .html template to push into layout
- The layout, login, and apology pages are already created.

#### Specification
1. Complete the implementation of a `register` page allowing the user to create a new record in the users table
    - Require that a user input a username with the text field whose `name` is `username`
    - Require the user input a password with the text field whose `name` is `password`
    - Require the user input a confirmation password with the text field whose `name` is `confirmation`
    - Submit the input via `POST` to `/register`
      - Return an apology if the username already exists in the db, the password and confirmation do not match, or
      a password is missing
    - `INSERT` the new user into `users` with a hashed password. Hash the password with the `generate_password_hash`
   function
    - On successful registration, route the user to the login page
2. Complete the implementation of a `quote` page that allows a user to look up a stock's current price
    - Require that a user input a stock's symbol as a text field whose `name` is `symbol`
    - Submit the input via `POST` to `/quote`
    - You'll want to create a `quote.html` template rendered when `/quote` request id done via `GET`, and a 
   `quoted.html` rendered when the quote is complete which embeds the values from `lookup`
3. Complete the implementation of a `buy` page that enables the user to purchase stocks
    - Require that a user input a stock's symbol, via an input whose `name` is `symbol`
      - Render an apology if the symbol does not exist (as per the return value of `lookup`)
    - Require that a user input the number of shares, via a text field whose `name` is `shares`
      - Render an apology if the input is not a positive number
    - Submit the user's input via `POST` to `/buy`
    - On completion, redirect the user to the homepage
    - Add one or more tables to `finance.db` to keep track of the purchase. Store enough information so you know who
   bought what at, at what price, and when it was purchased
      - Use appropriate SQLite types with `UNIQUE` and non-`UNIQUE` fields
    - Render an apology without completing a purchase if the user cannot afford the purchase
4. Complete the implementation of `index` that displays an HTML table, summarizing for the currently logged in user,
which stocks the user owns, the number of shares, the current price of each stock, and the total value of each holding.
Also display the user's current cash balance and a grand total (cash + stocks' total value)
    - Tip: Use `lookup` for each stock to get the current price
    - Tip: Likely will need a combination of `GROUP BY`, `HAVING`, `SUM`, and `WHERE`
5. Complete the implementation of `sell` page that enables the user to shell shares of a stock (that they own)
    - Require that a user input a stock's symbol, via a select menu who `name` is `symbol`
      - Render an apology if the user fails to select a stock, or the user does not own any shares in that stock
    - Require that a user input a number of shares to sell, via an input whose `name` is `shares`
      - Render an apology if the input is not a positive integer, or the user does not own that many shares
    - Submit the user's input via `POST` to `/sell`
    - Upon completion, redirect the user to the home page
6. Complete the implementation of `history` page that displays an HTML table to display the complete history
of a users transactions, listing row by row each and every buy and every sell
    - For each row, make clear whether a stock was bought or sold, including the symbol, the purchase or sale price,
   the number of shares bought or sold, and the date and time the transaction occurred
    - You may need to adjust the table created during implementation of the `buy` section
7. Optional
    - Allow users to change their password
    - Allow user to add additional cash to their accounts
    - Allow users to buy or sell shares they own via the `index` page without having to input the symbol