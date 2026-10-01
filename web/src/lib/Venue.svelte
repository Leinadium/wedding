<script lang="ts">
  import { fade } from "svelte/transition";
  import Text from "./text/Text.svelte";

  let {
    closeCb,
  }: {
    closeCb: () => void;
  } = $props();
</script>

<div class="venue-wrapper" transition:fade={{ duration: 200 }}>
  <div class="closable">
    <button class="close" onclick={closeCb}>&times;</button>
    <div class="venue-container">
      <div class="parents formal">
        <p class="middle"><Text key="venue-parentsmiddle" /></p>
        <div class="parents-side">
          <p class="left"><Text key="venue-parentsleft" /></p>
          <p class="right"><Text key="venue-parentsright" /></p>
        </div>
      </div>

      <span class="name cursive">
        <Text key="venue-name" />
      </span>

      <span class="description formal">
        <Text key="venue-description" />
      </span>

      <div class="actions">
        <div class="action">
          <span class="action-description formal">
            <Text key="venue-calendardescription" />
          </span>
        </div>

        <div class="action">
          <span class="action-description formal">
            <Text key="venue-locationdescription" />
          </span>
        </div>
      </div>

      <div class="footer">
        <div class="footer-line"></div>
        <p class="footer-text formal">
          <Text key="venue-footer" />
        </p>
      </div>
    </div>
  </div>
</div>

<style>
  .venue-wrapper {
    position: fixed;
    top: 0;
    left: 0;

    width: 100vw;
    height: 100vh;

    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
    align-items: center;

    background-color: rgba(0, 0, 0, 0.6);
    z-index: 999;
  }

  .closable {
    width: calc(80vw + 2rem);
    max-width: calc(800px + 2rem);
    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
    align-items: center;
    padding: 1rem;
  }

  .venue-container {
    box-sizing: border-box;
    position: relative;
    width: 100%;
    max-width: 800px;
    max-height: 90vh;
    overflow-y: auto;

    border-radius: 8px;
    padding: 1rem 1rem 1rem 1rem;
    box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.3);

    background-image: url("../assets/venue/background.png");
    background-position: center;
    background-size: cover;

    display: flex;
    flex-flow: column nowrap;
    justify-content: start;
    align-items: center;
  }

  .close {
    align-self: flex-end;
    background: none;
    border: none;
    font-size: 2rem;
    color: #dedacdff;
    cursor: pointer;
    z-index: 10;

    border-radius: 1rem;
    border: 1px solid transparent;
  }

  .parents {
    width: 100%;
    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
    align-items: center;

    color: #dedacda0;
    font-style: italic;
  }

  .parents-side {
    display: flex;
    width: 90%;
    flex-flow: row nowrap;
    justify-content: space-between;
  }

  .left {
    text-align: left;
  }

  .middle {
    text-align: center;
  }

  .right {
    text-align: right;
  }

  .name {
    text-align: center;
    color: #dedacdff;
  }

  .description {
    font-style: italic;
    text-align: center;

    color: #dedacdc0;
  }

  .actions {
    width: 100%;
    display: flex;
    align-items: center;
  }

  .action {
    margin: 0;
    border: 0;
    padding: 0;
    background: transparent;
    text-decoration: none;
    flex: 1 1 0px;
    box-sizing: border-box;

    text-align: center;
    color: #dedacdff;

    display: flex;
    flex-flow: column nowrap;
    align-items: center;

    border: 1px solid transparent;
    border-radius: 1rem;
  }

  button {
    cursor: pointer;
  }

  button:hover {
    border: 1px solid #dedacdff;
  }

  .footer {
    display: flex;
    flex-flow: column;
    justify-content: start;
    align-items: center;
  }

  .footer-line {
    border-top: 0.1rem solid #dedacdff;
  }

  .footer-text {
    text-align: center;
    font-style: italic;
    color: #dedacda0;
  }

  p {
    margin: 0;
  }

  @media only screen and (max-width: 768px) {
    .parents {
      height: 5rem;
      font-size: 1.3rem;
      gap: 0.5rem;
    }
    .parents-side {
      width: 100%;
    }

    .name {
      font-size: 6rem;
      line-height: 5rem;
      margin-top: 3rem;
    }
    .description {
      font-size: 1.2rem;
      margin-bottom: 2rem;
    }
    .actions {
      flex-flow: column nowrap;
      justify-content: center;
      align-items: center;
    }
    .action {
      width: 100%;
      padding: 1rem;
      flex-flow: row nowrap;
      justify-content: center;
      gap: 1.5rem;
    }
    .action-description {
      font-size: 1.5rem;
    }
    .footer {
      margin-top: 1.2rem;
    }
    .footer-line {
      width: 3rem;
      height: 0.5rem;
    }
    .footer-text {
      font-size: 1.3rem;
    }
  }

  @media only screen and (min-width: 768px) {
    .parents {
      height: 4rem;
      font-size: 1.5rem;
    }
    .name {
      font-size: 6rem;
      line-height: 5rem;
      margin-top: 3rem;
    }
    .description {
      font-size: 1.2rem;
      margin-bottom: 4rem;
    }
    .actions {
      height: 8rem;
      flex-flow: row nowrap;
      justify-content: space-around;
    }
    .action {
      width: 30%;
      padding: 1rem;
    }
    .action-description {
      font-size: 1.5rem;
    }
    .footer {
      margin-top: 3rem;
    }
    .footer-line {
      width: 3rem;
      height: 0.5rem;
    }
    .footer-text {
      font-size: 1.4rem;
    }
  }
</style>
