import React from 'react';
import FormComponent from './components/FormComponent';
import TableComponent from './components/TableComponent';

const App = () => {
    return (
        <div>
            <h1>Website Information Manager</h1>
            <FormComponent />
            <TableComponent />
        </div>
    );
};

export default App;