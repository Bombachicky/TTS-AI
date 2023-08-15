// components/Navbar.js
'use client'
import React from 'react';
import {FaBars,FaTimes} from "react-icons/fa";
import Link from 'next/link';
import { useRef } from 'react';



const Navbar = () =>{

    const navRef = useRef<HTMLElement>(null);
    
    const showNavbar = () => {
        navRef.current?.classList.toggle("responsive_nav")
    }
    return(<>
        <header className="text-primary flex items-center justify-between h-20 pt-0 px-8 bg-secondary">
            <Link href="/" >
                <h3 className='font-semibold'>OverTone</h3>
            </Link>
            <nav ref={navRef}>
                <Link href="/sign-in" className='nav-link'>
                    Sign In
                </Link>
                <Link href="/sign-up" className='nav-link'>
                    Sign Up
                </Link>
                <Link href="/" className='nav-link'>
                    Donate IM POOR
                </Link>
                <button  className="nav-btn nav-close-btn text-btn" onClick={showNavbar} >
                    <FaTimes/>
                </button>
            </nav>
            <button className="nav-btn nav-close-btn  text-btn" onClick={showNavbar}>
                <FaBars/>
            </button>
        </header>
    </>)}



// const Navbar = () => {
//   return (
//     <nav className="bg-secondary p-4">
//       <div className="max-w-screen-xl mx-auto">
//         <div className="flex justify-between items-center">
//           <div className="text-white font-bold">OverTone</div>
//           <ul className="flex space-x-4">
            
//             <li>
//               <a className="text-white" href="#">About</a>
//             </li>
//             <li>
//               <a className="text-white" href="#">Services</a>
//             </li>
//             <li>
//               <a className="text-white" href="#">Contact</a>
//             </li>
//           </ul>
//         </div>
//       </div>
//     </nav>
//   );
// };

export default Navbar;