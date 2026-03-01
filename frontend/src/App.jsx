import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import QuestionList from './pages/QuestionList';
import AddQuestion from './pages/AddQuestion';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<QuestionList />} />
        <Route path="/add" element={<AddQuestion />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;