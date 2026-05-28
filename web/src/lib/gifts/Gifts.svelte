<script lang="ts">
  import { fade } from "svelte/transition";
  import { api, type Product } from "../api";

  // State for products and the selected product for the popup
  let products = $state<Product[]>([]);
  let selectedProduct = $state<Product | null>(null);
  let loadingPayment = $state(false);

  // Fetch products on mount
  $effect(() => {
    // products = [
    //   {
    //     id: "123",
    //     name: "produto 1",
    //     imageURL: "",
    //     priceBRL: 1234,
    //     purchased: false,
    //   },
    //   {
    //     id: "123",
    //     name: "produto 1",
    //     imageURL: "",
    //     priceBRL: 1234,
    //     purchased: false,
    //   },
    //   {
    //     id: "123",
    //     name: "produto 1",
    //     imageURL: "",
    //     priceBRL: 1234,
    //     purchased: false,
    //   },
    //   {
    //     id: "123",
    //     name: "produto 1",
    //     imageURL: "",
    //     priceBRL: 1234,
    //     purchased: false,
    //   },
    // ];
    api.getProducts().then((res) => (products = res.products));
  });

  function formatPrice(price: number): string {
    if (isNaN(price) || price === 0) return "--";
    return `R$ ${Math.trunc(price / 100)},${price % 100}`;
  }

  async function handlePurchase(id: string) {
    loadingPayment = true;
    try {
      const { payment } = await api.getPaymentUrl(id);
      window.open(payment.url, "_blank");
    } finally {
      loadingPayment = false;
    }
  }
</script>

<div class="gallery">
  {#each products as product}
    <button
      class="product-box {product.purchased ? 'purchased' : ''}"
      onclick={() => (selectedProduct = product)}
    >
      <img src={product.imageUrl} alt={product.name} />
      <div class="info">
        <h3>{product.name}</h3>
        <p>{formatPrice(product.priceBrl)}</p>
      </div>
      {#if product.purchased}
        <span class="status-tag">Presenteado!</span>
      {/if}
    </button>
  {/each}
</div>

{#if selectedProduct}
  <div
    class="modal-backdrop"
    transition:fade
    onclick={() => (selectedProduct = null)}
  >
    <div class="modal-card" onclick={(e) => e.stopPropagation()}>
      <button class="close-btn" onclick={() => (selectedProduct = null)}
        >&times;</button
      >

      <img src={selectedProduct.imageUrl} alt={selectedProduct.name} />
      <h2>{selectedProduct.name}</h2>
      <p class="price">{formatPrice(selectedProduct.priceBrl)}</p>

      {#if selectedProduct.purchased}
        <p class="purchased-msg">
          This gift has already been purchased. Thank you!
        </p>
      {:else}
        <button
          class="buy-btn"
          disabled={loadingPayment}
          onclick={() => handlePurchase(selectedProduct!.id)}
        >
          {loadingPayment ? "Generating Link..." : "Give this Gift"}
        </button>
      {/if}
    </div>
  </div>
{/if}

<style>
  /* 1. The Grid Container */
  .gallery {
    display: grid;
    /* 2 columns on mobile */
    grid-template-columns: repeat(2, 1fr);
    gap: 12px;
    padding: 1rem;
    width: 100%;
    box-sizing: border-box;
  }

  /* 4 columns on desktop/tablets */
  @media (min-width: 768px) {
    .gallery {
      grid-template-columns: repeat(4, 1fr);
      gap: 20px;
    }
  }

  /* 2. The Square Product Card */
  .product-box {
    position: relative;
    aspect-ratio: 1 / 1; /* Forces perfect square */
    width: 100%;
    display: flex;
    flex-direction: column;
    background: #ffffff;
    border: 1px solid #eaeaea;
    border-radius: 8px; /* Small rounded corners */
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

  /* Image handling within the square */
  .product-box img {
    width: 100%;
    flex-grow: 1;
    min-height: 0; /* Important for flex-squish behavior */
    object-fit: cover;
  }

  /* Bottom text area */
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
    text-overflow: ellipsis; /* Keeps layout tight */
  }

  .info p {
    font-size: 0.85rem;
    font-weight: bold;
    margin: 4px 0 0 0;
    color: #8a7b6e;
  }

  /* 3. Purchased Logic */
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

  /* 4. Modal / Popup Styles */
  .modal-backdrop {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.6);
    backdrop-filter: blur(2px);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 1000;
    padding: 20px;
  }

  .modal-card {
    background: white;
    width: 100%;
    max-width: 350px;
    border-radius: 16px;
    padding: 24px;
    position: relative;
    text-align: center;
    box-shadow: 0 10px 25px rgba(0, 0, 0, 0.2);
  }

  .modal-card img {
    width: 100%;
    aspect-ratio: 1;
    object-fit: cover;
    border-radius: 12px;
    margin-bottom: 16px;
  }

  .modal-card h2 {
    font-size: 1.25rem;
    margin: 0 0 8px 0;
  }

  .price-large {
    font-size: 1.1rem;
    color: #8a7b6e;
    font-weight: bold;
    margin-bottom: 20px;
  }

  .buy-btn {
    background: #8a7b6e;
    color: white;
    border: none;
    padding: 14px;
    border-radius: 30px;
    font-size: 1rem;
    font-weight: bold;
    width: 100%;
    cursor: pointer;
    transition: background 0.2s;
  }

  .buy-btn:active {
    background: #6f6258;
  }

  .buy-btn:disabled {
    background: #ccc;
    cursor: not-allowed;
  }

  .close-btn {
    position: absolute;
    top: 12px;
    right: 12px;
    background: #eee;
    border: none;
    width: 30px;
    height: 30px;
    border-radius: 50%;
    font-size: 1.2rem;
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
  }
</style>
