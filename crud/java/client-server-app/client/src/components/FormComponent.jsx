import React, { useState } from 'react';

const FormComponent = ({ onSubmit }) => {
    const [websiteName, setWebsiteName] = useState('');
    const [websiteLink, setWebsiteLink] = useState('');
    const [description, setDescription] = useState('');
    const [tags, setTags] = useState('');
    const [category, setCategory] = useState('');

    const handleSubmit = (e) => {
        e.preventDefault();
        onSubmit({ websiteName, websiteLink, description, tags, category });
        setWebsiteName('');
        setWebsiteLink('');
        setDescription('');
        setTags('');
        setCategory('');
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
                <select value={tags} onChange={(e) => setTags(e.target.value)}>
                    <option value="">Select Tag</option>
                    {/* Add options dynamically */}
                </select>
            </div>
            <div>
                <label>Category:</label>
                <select value={category} onChange={(e) => setCategory(e.target.value)}>
                    <option value="">Select Category</option>
                    {/* Add options dynamically */}
                </select>
            </div>
            <button type="submit">Submit</button>
        </form>
    );
};

export default FormComponent;