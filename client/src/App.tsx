import { useRef } from 'react'
import './App.css'
import { useUser } from './hooks/useUser'

function App() {
  const emailRef = useRef(null)
  const passwordRef = useRef(null)
  const user = useUser()

  const submitForm = async (isLogin: boolean = false) => {
    const emailRx = /^[\w.%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$/

    const email = emailRef.current.value
    const password = passwordRef.current.value
    
    console.log(import.meta.env.API_URL);
    
    if (emailRx.test(email)) {
      if (isLogin) {

        const resp = await fetch("/api/login", {
          method: "POST",
          credentials: 'include',
          headers: {
            "Content-Type": "application/json"
          },
          body: JSON.stringify({
            email,
            password
          })
        })
  
        console.log(resp.status);
  
      } else {
        const resp = await fetch("/api/user-account", {
          method: "POST",
          credentials: 'include',
          headers: {
            "Content-Type": "application/json"
          },
          body: JSON.stringify({
            email,
            password
          })
        })
  
        console.log(resp.status);
  
      }


    } else {

    }
  
    console.log(email)
    console.log(password)
  }

  const logout = async () => {
    const resp = await fetch("/api/logout", {
      method:"POST",
      headers:{"Content-Type": "application/json"}
    })

    console.log(resp.status)
  }



  return (
    <>
      <div className='w-[100vw] min-h-screen bg-white flex flex-col text-black'>
        <p>Email</p>
        <input type="email" ref={emailRef} className=' border' />

        <p>Password</p>
        <input type="password" ref={passwordRef} className=' border' />

        <button onClick={() => submitForm(false)} className='bg-white text-black border'>
          Submit
        </button>

        <button onClick={() => submitForm(true)} className='bg-white text-black border'>
          Login
        </button>

        <button onClick={logout} className='bg-white text-black border'>Logout</button>
      </div>
    </>
  )
}

export default App
