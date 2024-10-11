"use client";

import { useState } from 'react';
import { useRouter } from 'next/navigation';

const Login = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const router = useRouter();

  const handleSubmit = (e) => {
      e.preventDefault();
      // Handle login logic here
      console.log('Email:', email);
      console.log('Password:', password);
      router.push('/surveys'); 
  };

  return (
      <div>
          <form onSubmit={handleSubmit} className="flex flex-col">
              <label>
                  Email:
                  <input
                      type="email"
                      value={email}
                      onChange={(e) => setEmail(e.target.value)}
                      required
                  />
              </label>
              <label>
                  Password:
                  <input
                      type="password"
                      value={password}
                      onChange={(e) => setPassword(e.target.value)}
                      required
                  />
              </label>
              <button className="btn" type="submit" link>Login</button>
          </form>
      </div>
  );
};

export default function Home() {
  return (
    <div className="p-10 flex flex-col justify-center items-center h-screen">
      <h1 className="text-center">Enov8tive Feedback System</h1>
      <Login />
    </div>
  );
}
