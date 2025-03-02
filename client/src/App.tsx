import { useRef } from 'react'
import './App.css'

function App() {
  const emailRef = useRef(null)
  const passwordRef = useRef(null)

  const submitForm = async () => {
    const emailRx = /^[\w.%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$/

    const email = emailRef.current.value
    const password = passwordRef.current.value
    
    console.log(import.meta.env.API_URL);
    

    if (emailRx.test(email)) {
      const resp = await fetch("https://lalocura-go-production.up.railway.app" + "/user-account", {
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

    }
  
    console.log(email)
    console.log(password)
  }

  return (
    <>
      <div className='w-[100vw] min-h-screen bg-white flex flex-col text-black'>
        <p>Email</p>
        <input type="email" ref={emailRef} className=' border' />

        <p>Password</p>
        <input type="password" ref={passwordRef} className=' border' />

        <button onClick={submitForm} className='bg-white text-black border'>
          Submit
        </button>
      </div>
    </>
  )
}

export default App
