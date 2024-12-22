import React, { useEffect, useState } from 'react';

const DataViewer: React.FC = () => {
    const [data, setData] = useState<any[]>([]);
    const [loading, setLoading] = useState<boolean>(true);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch('/api/data'); // Adjust the API endpoint as needed
                const result = await response.json();
                setData(result);
            } catch (error) {
                console.error('Error fetching data:', error);
            } finally {
                setLoading(false);
            }
        };

        fetchData();
    }, []);

    if (loading) {
        return <div>Loading...</div>;
    }

    return (
        <div>
            <h2>Data Viewer</h2>
            <table>
                <thead>
                    <tr>
                        <th>Website Name</th>
                        <th>Link</th>
                        <th>Description</th>
                        <th>Tags</th>
                        <th>Categories</th>
                    </tr>
                </thead>
                <tbody>
                    {data.map((item) => (
                        <tr key={item.id}>
                            <td>{item.websiteName}</td>
                            <td>{item.link}</td>
                            <td>{item.description}</td>
                            <td>{item.tags.join(', ')}</td>
                            <td>{item.categories.join(', ')}</td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
};

export default DataViewer;