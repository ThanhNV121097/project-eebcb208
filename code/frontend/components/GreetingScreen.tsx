import styles from './GreetingScreen.module.css';

type GreetingResponse = {
  text: string;
};

const apiOrigin = process.env.API_ORIGIN ?? 'http://backend:8080';

export default async function GreetingScreen() {
  const response = await fetch(`${apiOrigin}/v1/greeting`, { cache: 'no-store' });
  const body = (await response.json()) as GreetingResponse;

  return (
    <main className={styles.screen} aria-label="Greeting screen">
      <h1 className={styles.greeting}>{body.text}</h1>
      <div className={styles.hint} aria-hidden="true">
        White screen · centered black text
      </div>
    </main>
  );
}
