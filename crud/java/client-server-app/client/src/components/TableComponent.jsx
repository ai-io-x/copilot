import React, { useEffect, useState } from 'react';

const TableComponent = () => {
    const [data, setData] = useState([]);
    const [searchTerm, setSearchTerm] = useState('');
    const [sortField, setSortField] = useState('websiteName');

    useEffect(() => {
        fetchData();
    }, []);

    const fetchData = async () => {
        const response = await fetch('http://localhost:8080/api/data');
        const result = await response.json();
        setData(result);
    };

    const handleDelete = async (id) => {
        await fetch(`http://localhost:8080/api/data/${id}`, {
            method: 'DELETE',
        });
        fetchData();
    };

    const filteredData = data.filter(item =>
        item.websiteName.toLowerCase().includes(searchTerm.toLowerCase()) ||
        item.category.toLowerCase().includes(searchTerm.toLowerCase()) ||
        item.tags.some(tag => tag.toLowerCase().includes(searchTerm.toLowerCase()))
    );

    const sortedData = filteredData.sort((a, b) => {
        if (a[sortField] < b[sortField]) return -1;
        if (a[sortField] > b[sortField]) return 1;
        return 0;
    });

    return (
        <div>
            <input
                type="text"
                placeholder="Search..."
                value={searchTerm}
                onChange={(e) => setSearchTerm(e.target.value)}
            />
            <table>
                <thead>
                    <tr>
                        <th onClick={() => setSortField('websiteName')}>Website Name</th>
                        <th onClick={() => setSortField('link')}>Link</th>
                        <th onClick={() => setSortField('description')}>Description</th>
                        <th onClick={() => setSortField('category')}>Category</th>
                        <th onClick={() => setSortField('tags')}>Tags</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {sortedData.map(item => (
                        <tr key={item.id}>
                            <td>{item.websiteName}</td>
                            <td>{item.link}</td>
                            <td>{item.description}</td>
                            <td>{item.category}</td>
                            <td>{item.tags.join(', ')}</td>
                            <td>
                                <button onClick={() => handleDelete(item.id)}>Delete</button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
};

export default TableComponent;