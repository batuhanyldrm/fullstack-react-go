import './App.css';
import Blog from './components/Blog';
import React from 'react';
import InputField from './components/InputField';
function App() {
  return (
    <div className="App">
      <h1 style={{display:"flex", justifyContent:"center"}}>blog</h1>
      <Blog />
      <InputField/>
    </div>
  );
}

export default App;
