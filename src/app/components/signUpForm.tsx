"use client"
import axios from "axios"; // Import the axios library
import React, { useState } from 'react';
import { CognitoUserPool, CognitoUserAttribute, CognitoUser } from 'amazon-cognito-identity-js';



const poolData = {
  UserPoolId: 'us-east-1_GJv3BEuQQ',
  ClientId: '4ab88mrp1nq1om903il5lnmerv'
};

const userPool = new CognitoUserPool(poolData);



export default function SignUp() {

  const [email, setEmail] = useState('');
  const [userName, setUserName] = useState('');
  const [isSignedUp, setIsSignedUp] = useState(false);
  const [isCodeSent, setIsCodeSent] = useState(false);

  // Handle Sign Up Form
  const handleSignUp = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    // Collect user info
    const formData = new FormData(e.currentTarget);
    const email = formData.get("email") as string;
    const password = formData.get("password") as string;
    const username = formData.get("username") as string;


    // Create a new attribute list for Cognito user
    const attributeList = [
      new CognitoUserAttribute({ Name: 'email', Value: email }),
      // Add any other attributes you've configured in your user pool
    ];

    // Sign up User and send verification code
    userPool.signUp(email, password, attributeList, [], (err, result) => {
      if (err) {
        console.error('Error signing up:', err);
        return;
      }
      // If you want to use your API gateway after successfully registering with Cognito
      // you can make an axios request here.

      console.log('User signed up:', result);
      setEmail(email);
      setUserName(username);
      setIsCodeSent(true);

    });




  };

  // Handle Code Verification 
  const handleCodeVerification = () => {
    const verificationInputElement = document.getElementById('verification-code') as HTMLInputElement | null;

    if (!verificationInputElement) {
      console.error('Verification code input not found');
      return;
    }

    const verificationCode = verificationInputElement.value;

    const cognitoUser = new CognitoUser({
      Username: email,
      Pool: userPool,
    });

    cognitoUser.confirmRegistration(verificationCode, true, async (err, result) => {
      if (err) {
        console.error('Error verifying code:', err);
        return;
      }
      console.log('Code verified:', result);
      setIsSignedUp(true);

      // Now, after successful code verification, add the user to the DB.
      // Assuming userData is available in this context. If not, you might need
      // to maintain userData in the state.
      const userData = {
        email: email,
        username: userName,
        speed: 0,
        pitch: 0,
      };
      try {
        const response = await axios.post(
          'https://5f0ek1er9i.execute-api.us-east-1.amazonaws.com/prod/users/sign-up',
          userData
        );
        console.log('User added to DB:', response.data);
        // Handle other success logic here, like redirecting the user or showing a confirmation.
      } catch (error) {
        console.error('Error adding to DB:', error);
      }
    });




  };



  return (
    <div className="flex min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8">
      <div className="sm:mx-auto sm:w-full sm:max-w-sm">
        <h3 className="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-primary">PENDING LOGO</h3>

        {isSignedUp ? (
          <div className="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-primary">
            Thank you for signing up!
          </div>
        ) : isCodeSent ? (
          <div className="flex items-center gap-1">
            <input className="block rounded-md border-0 py-1.5 text-formText shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-btn sm:text-sm sm:leading-6" id="verification-code" type="text" placeholder="Enter verification code" />
            <button
              onClick={handleCodeVerification}
              className=" justify-center rounded-md bg-btn px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-btn focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
              id="verification-code"
            >
              Submit
            </button>
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