'use client';
import Image from 'next/image';
import GoogleMapComponent from './components/GoogleMap';

export default function Home() {
  return (
    <div>
      <h1>Image Splitter</h1>
      <input type="file" accept="image/*" />
      <Image src="/sample.jpg" width="100" height="100" alt="img" />
      <GoogleMapComponent />
    </div>
  );
}
