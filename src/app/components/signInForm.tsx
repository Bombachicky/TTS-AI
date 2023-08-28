"use client"
import { CognitoUserPool } from 'amazon-cognito-identity-js';
import Navbar from "../components/navbar"
import React, { useState } from "react";
import Cookies from 'js-cookie'

import { CognitoUser, CognitoUserAttribute, AuthenticationDetails } from 'amazon-cognito-identity-js';

import axios from "axios"; // Import the axios library

const poolData = {
  UserPoolId: 'us-east-1_GJv3BEuQQ',
  ClientId: '4ab88mrp1nq1om903il5lnmerv'
};
const userPool = new CognitoUserPool(poolData);

export default function SignIn() {
  const [isSignedIn, setSignedIn] = useState(false);


  const handleSignIn = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    // Collect form information
    const formData = new FormData(e.currentTarget);
    const email = formData.get("email") as string;
    const password = formData.get("password") as string;



    const cognitoUser = new CognitoUser({
      Username: email,
      Pool: userPool,
    });

    // Authenticate user on sign in
    const authDetails = new AuthenticationDetails({
      Username: email,
      Password: password
    });

    // Authenticate user fucntion that initaties at handle sign in
    cognitoUser.authenticateUser(authDetails, {
      onSuccess: data => {
        console.log('onSuccess:', data);
        setSignedIn(true);

        Cookies.set('userToken', data.getIdToken().getJwtToken());
      },
      onFailure: err => {
        console.error('onFailure:', err);
      },
      newPasswordRequired: data => {
        console.log('newPasswordRequired:', data);
      }
    });


  };


  //TODO FIX IS SIGNED IN 
  return (<>

    {isSignedIn ? (
      <div className="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-primary">
        Thank you for signing in!
      </div>
    ) :
      (
        <div className="flex min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8">
          <div className="sm:mx-auto sm:w-full sm:max-w-sm">

            <h3 className="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-primary">PENDING LOGO</h3>
            <h2 className="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-primary">
              Sign in to your account
            </h2>
          </div>

          <div className="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
            {/* Changing form post to get */}
            <form className="space-y-6" onSubmit={handleSignIn}>
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
                  <div className="text-sm">
                    <a href="#" className="font-semibold text-btn hover:text-btn">
                      Forgot password?
                    </a>
                  </div>
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
                <button
                  type="submit"
                  className="flex w-full justify-center rounded-md bg-btn px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-btn focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                >
                  Sign in
                </button>
              </div>
            </form>

            <p className="mt-10 text-center text-sm text-gray-500">
              Not a member?{' '}
              <a href="#" className="font-semibold leading-6 text-btn hover:text-btn">
                Start a 14 day free trial
              </a>
            </p>
          </div>
        </div>)}</>);
}