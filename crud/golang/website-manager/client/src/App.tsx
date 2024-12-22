import React, { useState, useEffect } from 'react';
import WebsiteForm from './components/WebsiteForm';
import WebsiteTable from './components/WebsiteTable';

const App = () => {
    const [websites, setWebsites] = useState([]);
    const [loading, setLoading] = useState(true);

    const fetchWebsites = async () => {
        const response = await fetch('http://localhost:8080/websites');
        const data = await response.json();
        setWebsites(data);
        setLoading(false);
    };

    const addWebsite = async (website) => {
        await fetch('http://localhost:8080/websites', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(website),
        });
        fetchWebsites();
    };

    const updateWebsite = async (id, updatedWebsite) => {
        await fetch(`http://localhost:8080/websites/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(updatedWebsite),
        });
        fetchWebsites();
    };

    const deleteWebsite = async (id) => {
        await fetch(`http://localhost:8080/websites/${id}`, {
            method: 'DELETE',
        });
        fetchWebsites();
    };

    useEffect(() => {
        fetchWebsites();
    }, []);

    return (
        <div>
            <h1>Website Manager</h1>
            <WebsiteForm onAddWebsite={addWebsite} />
            {loading ? (
                <p>Loading...</p>
            ) : (
                <WebsiteTable
                    websites={websites}
                    onUpdateWebsite={updateWebsite}
                    onDeleteWebsite={deleteWebsite}
                />
            )}
        </div>
    );
};

export default App;