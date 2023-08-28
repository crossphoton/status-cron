'use client';
import Image from 'next/image'
import styles from './page.module.css'
import Button from '@/components/interaction/Button';

export default function Home() {
  return (
    <main className={styles.main}>
      <div className={styles.description}>
        <p>Status Cron by crossphoton</p>
        <div>
          <a
            href="https://vercel.com?utm_source=create-next-app&utm_medium=appdir-template&utm_campaign=create-next-app"
            target="_blank"
            rel="noopener noreferrer"
          >
            Powered by{" "}
            <Image
              src="/vercel.svg"
              alt="Vercel Logo"
              className={styles.vercelLogo}
              width={100}
              height={24}
              priority
            />
          </a>
        </div>
      </div>

      <div className={styles.center}>
        <Image
          className={styles.logo}
          src="/next.svg"
          alt="Next.js Logo"
          width={180}
          height={37}
          priority
        />
      </div>

      <div className={styles.center+" test"}>
        <a href="/dashboard">
          <Button text="Go to your dashboard" />
        </a>
      </div>

      <div className={styles.grid}>
        <a
          href="/readme"
          className={styles.card}
          target="_blank"
          rel="noopener noreferrer"
        >
          <h2>
            Docs <span>-&gt;</span>
          </h2>
          <p>Find in-depth information about Next.js features and API.</p>
        </a>

        <a
          href="/demo"
          className={styles.card}
          target="_blank"
          rel="noopener noreferrer"
        >
          <h2>
            Demo App <span>-&gt;</span>
          </h2>
          <p>Go through the demo app and explore the features.</p>
        </a>
      </div>
    </main>
  );
}
