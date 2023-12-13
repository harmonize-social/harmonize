<script>
    export let title = 'Panel';
  </script>
  
  <style>
    .panel {
      background-color: #B931FC;
      border-radius: 10px;
      padding: 20px;
      box-shadow: 0 0 5px rgba(0, 0, 0, 0.1);
      display: flex;
      flex-direction: column;
      align-items: center;
      height: 100vh;
    }
  
    .panel h2 {
      flex: 0 0 auto;
    }
  
    .panel slot {
      flex: 1 1 auto;
      overflow: auto;
    }
  </style>
  
  <div class="panel">
    <h2>{title}</h2>
    <slot></slot>
  </div>