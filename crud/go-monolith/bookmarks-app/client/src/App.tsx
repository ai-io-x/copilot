import React from 'react';
import BookmarkForm from './components/BookmarkForm';
import BookmarkTable from './components/BookmarkTable';

const App: React.FC = () => {
    return (
        <div>
            <h1>Bookmarks Manager</h1>
            <BookmarkForm />
            <BookmarkTable />
        </div>
    );
};

export default App;