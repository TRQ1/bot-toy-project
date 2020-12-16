import app from './index.js';

const PORT = 8080;

const handleListening = () =>
    console.log(`Listening on: http://localhost:${PORT}`);

app.listen(PORT, handleListening);
