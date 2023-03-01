import { useState } from 'react';
import { useRouter } from 'next/router';

export default function Login() {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [form, setForm] = useState({ roll: '', password: '' });
  const router = useRouter();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError(null);
    try {
      await login(form.roll, form.password);
      router.push('/home');
    } catch (error) {
      setError(error.message);
      setLoading(false);
    }
  };

  const handleChange = (e) => {
    setForm((prevForm) => ({ ...prevForm, [e.target.name]: e.target.value }));
  };

  return (
    <div>
      <h1>Login</h1>
      <form onSubmit={handleSubmit}>
        <div>
            <p>Enter your IITK Roll No.</p>
          {/* <label htmlFor="roll">Roll:</label> */}
          <input type="text" id="roll" name="roll" value={form.roll} onChange={handleChange} />
        </div>
        <div>
          <p>Enter your Password</p>
          {/* <label htmlFor="password">Password:</label> */}
          <input type="password" id="password" name="password" value={form.password} onChange={handleChange} />
        </div>
        {error && <p>{error}</p>}
        <button type="submit" disabled={loading}>
          {loading ? 'Loading...' : 'Login'}
        </button>
      </form>
    </div>
  );
}
