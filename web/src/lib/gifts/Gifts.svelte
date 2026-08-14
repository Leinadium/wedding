<script lang="ts">
  import { fade } from "svelte/transition";
  import { api, type Product } from "../api";
  import Text from "../text/Text.svelte";
  import { formatPrice } from "../common";
  import Gift from "./Gift.svelte";
  import { onMount } from "svelte";

  let {
    closeCb,
  }: {
    closeCb: () => void;
  } = $props();

  let products = $state<Product[]>([]);
  let selectedProduct = $state<Product | null>(null);

  async function loadProducts() {
    try {
      let res = await api.getProducts();
      products = res.products;
    } catch (e) {
      console.error(e);
    }
  }

  onMount(loadProducts);
</script>

<div class="gifts-wrapper" transition:fade={{ duration: 300 }}>
  <div class="gifts-container">
    <button class="close-main" onclick={closeCb}>&times;</button>

    <div class="description">
      <span class="description-title cursive">
        <Text key="gifts-title" />
      </span>
      <span class="description-body formal">
        <Text key="gifts-description" />
      </span>
      <span class="description-body formal">
        <Text key="gifts-description2" />
      </span>
    </div>

    <div class="gallery">
      {#each products as product}
        <button
          class="product-box {product.purchased || !product.active
            ? 'purchased'
            : ''}"
          onclick={() => (selectedProduct = product)}
        >
          <img src={product.imageUrl} alt={product.name} />
          <div class="info">
            <h3>{product.name}</h3>
            <p>{formatPrice(product.priceBrl)}</p>
          </div>
          {#if product.purchased}
            <span class="status-tag">Presenteado!</span>
          {:else if !product.active}
            <span class="status-tag">Indisponível!</span>
          {/if}
        </button>
      {/each}
    </div>
  </div>
  {#if selectedProduct}
    <Gift
      product={selectedProduct}
      close={() => {
        selectedProduct = null;
        loadProducts(); // trigger reload
      }}
    />
  {/if}
</div>

<style>
  .gifts-wrapper {
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

  .gallery {
    position: relative;
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    gap: 12px;
    padding: 1rem;
    width: 100%;
    box-sizing: border-box;
  }

  @media (min-width: 768px) {
    .gallery {
      grid-template-columns: repeat(4, 1fr);
      gap: 20px;
    }
  }

  .product-box {
    position: relative;
    aspect-ratio: 1 / 1;
    width: 100%;
    display: flex;
    flex-direction: column;
    background: #ffffff;
    border: 1px solid #eaeaea;
    border-radius: 8px;
    overflow: hidden;
    cursor: pointer;
    padding: 0;
    transition:
      transform 0.2s ease,
      box-shadow 0.2s ease;
    -webkit-tap-highlight-color: transparent;
  }

  .product-box:hover {
    transform: translateY(-4px);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
  }

  .product-box img {
    width: 100%;
    flex-grow: 1;
    min-height: 0;
    object-fit: cover;
  }

  .info {
    padding: 8px;
    background: #ffffff;
    flex-shrink: 0;
    text-align: center;
    border-top: 1px solid #f5f5f5;
  }

  .info h3 {
    font-size: 0.8rem;
    margin: 0;
    color: #333;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .info p {
    font-size: 0.85rem;
    font-weight: bold;
    margin: 4px 0 0 0;
    color: #8a7b6e;
  }

  .product-box.purchased {
    background-color: #f9f9f9;
    filter: grayscale(1);
    opacity: 0.6;
    cursor: default;
  }

  .status-tag {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%) rotate(-15deg);
    background: rgba(255, 255, 255, 0.9);
    color: #555;
    padding: 3px 8px;
    border: 2px solid #555;
    font-weight: bold;
    font-size: 0.7rem;
    text-transform: uppercase;
    letter-spacing: 1px;
    pointer-events: none;
    z-index: 2;
  }

  .gifts-container {
    position: relative;
    width: 90%;
    max-width: 800px;
    max-height: 90vh;
    overflow-y: auto;

    border-radius: 8px;
    padding: 3rem 1rem 1rem 1rem;
    box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.3);

    background-image: url("../../assets/invite/texture.png");
    background-color: #f0f0f0;
    box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.3);

    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
  }

  .close-main {
    position: absolute;
    top: 1rem;
    right: 1rem;
    background: none;
    border: none;
    font-size: 1.5rem;
    color: #777;
    cursor: pointer;
    z-index: 10;
  }

  .description {
    display: flex;
    flex-flow: column;
    justify-content: start;
    align-items: center;
  }

  .description-title {
    font-size: 1.5rem;
  }

  .description-body {
    font-size: 1.5rem;
  }
</style>
