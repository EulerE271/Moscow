import React, { useState } from 'react';
import './App.css';

function App() {
  const [email, setEmail] = useState('');
  const [message, setMessage] = useState('')

  const handleSubmit = async () => {
    try {
      const response = await fetch('http://localhost:8080/send-glossary', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ email })
      });

      const data = await response.json();
      setMessage(data.message);
    } catch (err) {
      setMessage("Error sending email");
    };
  };

  return (
    <div className="App ">
      <div className="">
        <h1 className="">Send Words</h1>
        <input 
          type="email"
          placeholder="Enter your email"
          value={email}
          onChange={e => setEmail(e.target.value)}
          className=''
        />
        <button 
          onClick={handleSubmit}
          className='w-full bg-blue-400 text-white p-2 rounded hover:bg-indigo-400'
        >
          Send Words
        </button>
        {message && (
          <p className="mt-4 text-gray-600">{message}</p>
        )}
      </div>
    </div>
  );
}

export default App;
