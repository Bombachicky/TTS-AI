"use client"
import axios from "axios"; // Import the axios library
import React, { useState } from 'react';

export default function SignUp(){

  const [isSignedUp, setIsSignedUp] = useState(false);
 const handleSignUp = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const formData = new FormData(e.currentTarget);
    const email = formData.get("email") as string;
    console.log(email);
    const password = formData.get("password") as string;
    console.log(password);
    const username = formData.get("username") as string;

    

    const userData = {
      email: email,
      password: password,
      username : username,
      speed: 0,
      pitch: 0,
    };

    console.log(userData);

    try {
      const response = await axios.post(
        'https://5f0ek1er9i.execute-api.us-east-1.amazonaws.com/prod/users',
        userData
      );

      // Handle success response
      console.log('User signed up:', response.data);
      setIsSignedUp(true);

      // Redirect or update app state as needed
    } catch (error) {
      // Handle error
      console.log("this is error" + error);
      console.error('Error signing up:', error);
    }
  };


    return(

    
    
    <div className="flex min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8">
        <div className="sm:mx-auto sm:w-full sm:max-w-sm">
          <h3 className="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-primary">PENDING LOGO</h3>
           
           {isSignedUp ? (
        <div className="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-primary">
          Thank you for signing up!
        </div>
      ) : (
        <>
          <h2 className="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-primary">
            Sign up
          </h2>


            <div className="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
          <form className="space-y-6" onSubmit={handleSignUp}>
            <div>
              <label htmlFor="email" className="block text-sm font-medium leading-6 text-primary">
                Email address
              </label>
              <div className="mt-2">
                <input
                  id="email"
                  name="email"
                  type="email"
                  autoComplete="email"
                  required
                  className="block w-full rounded-md border-0 py-1.5 text-formText shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-btn sm:text-sm sm:leading-6"
                />
              </div>
            </div>

            <div>
              <div className="flex items-center justify-between">
                <label htmlFor="password" className="block text-sm font-medium leading-6 text-primary">
                  Password
                </label>
               
              </div>
              <div className="mt-2">
                <input
                  id="password"
                  name="password"
                  type="password"
                  autoComplete="current-password"
                  required
                  className="block w-full rounded-md border-0 py-1.5 text-formText shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                />
              </div>
            </div>

            <div>
              <div className="flex items-center justify-between">
                <label htmlFor="username" className="block text-sm font-medium leading-6 text-primary">
                  Username
                </label>
               
              </div>
              <div className="mt-2">
                <input
                  id="username"
                  name="username"
                  type="text"
                  autoComplete="username"
                  required
                  className="block w-full rounded-md border-0 py-1.5 text-formText shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                />
              </div>
            </div>

            <div>
              <button
                type="submit"
                className="flex w-full justify-center rounded-md bg-btn px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-btn focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
              >
                Sign up
              </button>
            </div>
          </form>

          
        </div>

          </>
      )}
    </div>
  </div>






      )
}