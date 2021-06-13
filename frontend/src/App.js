import React, { useState } from 'react';

const App = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const loginAction = async (email, password) => {
    const payload = { email, password };
  
    const response = await fetch('http://localhost:8080/authenticate', {
      method: 'POST',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(payload)
    });
  }

  return (
    <div>
      
      <input onChange={e => setEmail(e.target.value)} type='text' placeholder='email' name='email'/>
      <input onChange={e => setPassword(e.target.value)} type='password' placeholder='password' name='password'/>

      <button onClick={_ => loginAction(email, password)}>Login</button>

    </div>
  );
}

export default App;
