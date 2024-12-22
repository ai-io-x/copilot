import React, { useEffect, useState } from 'react';

const BookmarkTable = () => {
    const [bookmarks, setBookmarks] = useState([]);
    const [searchTerm, setSearchTerm] = useState('');
    
    useEffect(() => {
        fetchBookmarks();
    }, []);

    const fetchBookmarks = async () => {
        const response = await fetch('/api/bookmarks');
        const data = await response.json();
        setBookmarks(data);
    };

    const handleDelete = async (id) => {
        await fetch(`/api/bookmarks/${id}`, { method: 'DELETE' });
        fetchBookmarks();
    };

    const filteredBookmarks = bookmarks.filter(bookmark =>
        bookmark.name.toLowerCase().includes(searchTerm.toLowerCase())
    );

    return (
        <div>
            <input
                type="text"
                placeholder="Search bookmarks"
                value={searchTerm}
                onChange={(e) => setSearchTerm(e.target.value)}
            />
            <table>
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Link</th>
                        <th>Description</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {filteredBookmarks.map(bookmark => (
                        <tr key={bookmark.id}>
                            <td>{bookmark.name}</td>
                            <td>{bookmark.link}</td>
                            <td>{bookmark.description}</td>
                            <td>
                                <button onClick={() => handleDelete(bookmark.id)}>Delete</button>
                                {/* Add update functionality here */}
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
};

export default BookmarkTable;