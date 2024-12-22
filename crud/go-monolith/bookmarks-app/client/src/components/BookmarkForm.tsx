import React, { useState } from 'react';

const BookmarkForm = () => {
    const [websiteName, setWebsiteName] = useState('');
    const [websiteLink, setWebsiteLink] = useState('');
    const [description, setDescription] = useState('');
    const [tags, setTags] = useState('');
    const [category, setCategory] = useState('');

    const handleSubmit = (e) => {
        e.preventDefault();
        // Logic to send data to the REST server
    };

    return (
        <form onSubmit={handleSubmit}>
            <div>
                <label>Website Name:</label>
                <input 
                    type="text" 
                    value={websiteName} 
                    onChange={(e) => setWebsiteName(e.target.value)} 
                    required 
                />
            </div>
            <div>
                <label>Website Link:</label>
                <input 
                    type="url" 
                    value={websiteLink} 
                    onChange={(e) => setWebsiteLink(e.target.value)} 
                    required 
                />
            </div>
            <div>
                <label>Description:</label>
                <textarea 
                    value={description} 
                    onChange={(e) => setDescription(e.target.value)} 
                />
            </div>
            <div>
                <label>Tags:</label>
                <input 
                    type="text" 
                    value={tags} 
                    onChange={(e) => setTags(e.target.value)} 
                />
            </div>
            <div>
                <label>Category:</label>
                <select 
                    value={category} 
                    onChange={(e) => setCategory(e.target.value)} 
                >
                    <option value="">Select a category</option>
                    {/* Add category options here */}
                </select>
            </div>
            <button type="submit">Add Bookmark</button>
        </form>
    );
};

export default BookmarkForm;