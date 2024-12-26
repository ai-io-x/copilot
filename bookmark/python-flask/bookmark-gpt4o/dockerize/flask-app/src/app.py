from flask import Flask, render_template, request, redirect, url_for
import sqlite3

app = Flask(__name__)

# Initialize the database
def init_db():
    conn = sqlite3.connect('bookmarks.db')
    cursor = conn.cursor()
    cursor.execute('''
        CREATE TABLE IF NOT EXISTS bookmarks (
            id INTEGER PRIMARY KEY,
            name TEXT NOT NULL,
            link TEXT NOT NULL,
            description TEXT
        )
    ''')
    conn.commit()
    conn.close()

@app.route('/')
def index():
    return render_template('index.html')

@app.route('/add', methods=['POST'])
def add_bookmark():
    name = request.form['name']
    link = request.form['link']
    description = request.form['description']
    
    conn = sqlite3.connect('bookmarks.db')
    cursor = conn.cursor()
    cursor.execute('INSERT INTO bookmarks (name, link, description) VALUES (?, ?, ?)', (name, link, description))
    conn.commit()
    conn.close()
    
    return redirect(url_for('index'))

if __name__ == '__main__':
    init_db()
    app.run(debug=True, host='0.0.0.0')