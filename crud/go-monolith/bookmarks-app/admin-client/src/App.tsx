import React from 'react';
import LoginForm from './components/LoginForm';
import SqlEditor from './components/SqlEditor';
import DataViewer from './components/DataViewer';

const App: React.FC = () => {
    return (
        <div>
            <h1>Admin Panel</h1>
            <LoginForm />
            <SqlEditor />
            <DataViewer />
        </div>
    );
};

export default App;