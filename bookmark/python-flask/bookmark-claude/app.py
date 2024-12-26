from flask import Flask, render_template, request, redirect, url_for, flash
from flask_sqlalchemy import SQLAlchemy
from datetime import datetime

app = Flask(__name__)
app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///bookmarks.db'
app.config['SECRET_KEY'] = 'your-secret-key'
db = SQLAlchemy(app)

class Bookmark(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(100), nullable=False)
    url = db.Column(db.String(200), nullable=False)
    description = db.Column(db.Text)
    created_at = db.Column(db.DateTime, default=datetime.utcnow)

@app.route('/')
def index():
    bookmarks = Bookmark.query.order_by(Bookmark.created_at.desc()).all()
    return render_template('index.html', bookmarks=bookmarks)

@app.route('/add', methods=['POST'])
def add_bookmark():
    name = request.form['name']
    url = request.form['url']
    description = request.form['description']
    
    bookmark = Bookmark(name=name, url=url, description=description)
    db.session.add(bookmark)
    db.session.commit()
    
    flash('Bookmark added successfully!', 'success')
    return redirect(url_for('index'))

@app.route('/edit/<int:id>', methods=['GET', 'POST'])
def edit_bookmark(id):
    bookmark = Bookmark.query.get_or_404(id)
    
    if request.method == 'POST':
        bookmark.name = request.form['name']
        bookmark.url = request.form['url']
        bookmark.description = request.form['description']
        db.session.commit()
        flash('Bookmark updated successfully!', 'success')
        return redirect(url_for('index'))
    
    return render_template('edit.html', bookmark=bookmark)

@app.route('/delete/<int:id>')
def delete_bookmark(id):
    bookmark = Bookmark.query.get_or_404(id)
    db.session.delete(bookmark)
    db.session.commit()
    flash('Bookmark deleted successfully!', 'success')
    return redirect(url_for('index'))

if __name__ == '__main__':
    with app.app_context():
        db.create_all()
    app.run(debug=True)