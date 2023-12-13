<script>
    export let buttonText = 'Click me';
    export let link = '#';
  </script>

  <style>
    button {
      background-color: blue;
      color: white;
      padding: 10px 20px;
      font-size: 18px;
      border: none;
      border-radius: 5px;
      cursor: pointer;
    }
  </style>

  <a href={link}>
    <button>{buttonText}</button>
  </a>
