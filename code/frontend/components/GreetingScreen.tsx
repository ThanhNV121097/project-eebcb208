'use client';

import { useEffect, useState } from 'react';
import styles from './GreetingScreen.module.css';

type GreetingResponse = {
  text: string;
};

const apiBase = process.env.NEXT_PUBLIC_API_URL ?? '/api';

export default function GreetingScreen() {
  const [text, setText] = useState('');

  useEffect(() => {
    const controller = new AbortController();

    fetch(`${apiBase}/v1/greeting`, { signal: controller.signal })
      .then((response) => response.json())
      .then((body: GreetingResponse) => setText(body.text))
      .catch(() => {});

    return () => controller.abort();
  }, []);

  return (
    <main className={styles.screen} aria-label="Greeting screen">
      <h1 className={styles.greeting}>{text}</h1>
      <div className={styles.hint} aria-hidden="true">
        White screen · centered black text
      </div>
    </main>
  );
}
