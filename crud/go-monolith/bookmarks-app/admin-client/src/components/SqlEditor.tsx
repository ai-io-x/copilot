import React, { useState } from 'react';

const SqlEditor: React.FC = () => {
    const [sqlCommand, setSqlCommand] = useState('');
    const [result, setResult] = useState<any>(null);

    const handleExecute = async () => {
        try {
            const response = await fetch('/api/sql', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ command: sqlCommand }),
            });
            const data = await response.json();
            setResult(data);
        } catch (error) {
            console.error('Error executing SQL command:', error);
        }
    };

    return (
        <div>
            <h2>SQL Editor</h2>
            <textarea
                value={sqlCommand}
                onChange={(e) => setSqlCommand(e.target.value)}
                rows={10}
                cols={50}
                placeholder="Enter your SQL command here"
            />
            <br />
            <button onClick={handleExecute}>Execute</button>
            {result && (
                <div>
                    <h3>Result:</h3>
                    <pre>{JSON.stringify(result, null, 2)}</pre>
                </div>
            )}
        </div>
    );
};

export default SqlEditor;