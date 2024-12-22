import React, { useEffect, useState } from 'react';

const WebsiteTable = () => {
    const [websites, setWebsites] = useState([]);
    const [searchTerm, setSearchTerm] = useState('');
    
    useEffect(() => {
        fetchWebsites();
    }, []);

    const fetchWebsites = async () => {
        const response = await fetch('/api/websites'); // Adjust the API endpoint as needed
        const data = await response.json();
        setWebsites(data);
    };

    const handleDelete = async (id) => {
        await fetch(`/api/websites/${id}`, { method: 'DELETE' });
        fetchWebsites();
    };

    const filteredWebsites = websites.filter(website =>
        website.name.toLowerCase().includes(searchTerm.toLowerCase())
    );

    return (
        <div>
            <input
                type="text"
                placeholder="Search by website name"
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
                    {filteredWebsites.map(website => (
                        <tr key={website.id}>
                            <td>{website.name}</td>
                            <td>{website.link}</td>
                            <td>{website.description}</td>
                            <td>
                                <button onClick={() => handleDelete(website.id)}>Delete</button>
                                {/* Add update functionality here */}
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
};

export default WebsiteTable;