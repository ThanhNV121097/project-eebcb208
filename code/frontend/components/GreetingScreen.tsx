import styles from './GreetingScreen.module.css';
import { greetingResponse } from '../lib/mock/display-database-greeting';

export default function GreetingScreen() {
  return (
    <main className={styles.screen} aria-label="Greeting screen">
      <h1 className={styles.greeting}>{greetingResponse.text}</h1>
      <div className={styles.hint} aria-hidden="true">
        White screen · centered black text
      </div>
    </main>
  );
}
