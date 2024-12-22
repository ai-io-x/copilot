import React, { useState } from 'react';

const WebsiteForm = ({ onSubmit }) => {
    const [websiteName, setWebsiteName] = useState('');
    const [websiteLink, setWebsiteLink] = useState('');
    const [description, setDescription] = useState('');
    const [tags, setTags] = useState([]);
    const [categories, setCategories] = useState([]);

    const handleSubmit = (e) => {
        e.preventDefault();
        if (!websiteName || !websiteLink) {
            alert('Website name and link are required.');
            return;
        }
        const websiteData = { websiteName, websiteLink, description, tags, categories };
        onSubmit(websiteData);
        setWebsiteName('');
        setWebsiteLink('');
        setDescription('');
        setTags([]);
        setCategories([]);
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
                <select multiple value={tags} onChange={(e) => setTags([...e.target.selectedOptions].map(option => option.value))}>
                    {/* Options for tags should be populated here */}
                </select>
            </div>
            <div>
                <label>Categories:</label>
                <select multiple value={categories} onChange={(e) => setCategories([...e.target.selectedOptions].map(option => option.value))}>
                    {/* Options for categories should be populated here */}
                </select>
            </div>
            <button type="submit">Submit</button>
        </form>
    );
};

export default WebsiteForm;