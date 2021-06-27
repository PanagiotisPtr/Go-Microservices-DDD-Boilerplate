import React, { useState } from 'react';

const Register = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [verificationPassword, setVerificationPassword] = useState('');

  const registerAction = async (email, password, verificationPassword) => {
    if (password !== verificationPassword) {
      alert("Passwords don't match");
      return;
    }

    const payload = { email, password };
    const response = await fetch('http://localhost:8080/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(payload)
    });

    if (response.status === 200) {
      alert("Successfully Registered");
    }
  }

  return (
    <div>
      <h1>Register</h1>
      <input onChange={e => setEmail(e.target.value)} type='text' placeholder='email' name='email'/>
      <input onChange={e => setPassword(e.target.value)} type='password' placeholder='password' name='password'/>
      <input onChange={e => setVerificationPassword(e.target.value)} type='password' placeholder='password' name='password'/>

      <button onClick={_ => registerAction(email, password, verificationPassword)}>Register</button>

    </div>
  );
}

export default Register;