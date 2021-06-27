import React, { useState } from 'react';

const Login = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const loginAction = async (email, password) => {
    const payload = { email, password };
  
    const response = await fetch('http://localhost:8080/authenticate', {
      method: 'POST',
      credentials: 'include',
      mode: 'no-cors',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(payload)
    });

    console.log(response);
  }

  return (
    <div>
      <h1>Login</h1>
      <input onChange={e => setEmail(e.target.value)} type='text' placeholder='email' name='email'/>
      <input onChange={e => setPassword(e.target.value)} type='password' placeholder='password' name='password'/>

      <button onClick={_ => loginAction(email, password)}>Login</button>

    </div>
  );
}

export default Login;