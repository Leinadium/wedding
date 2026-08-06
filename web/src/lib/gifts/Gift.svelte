<script lang="ts">
  import { fade } from "svelte/transition";
  import { api, type Product, type Purchase } from "../api";
  import { formatPrice } from "../common";
  import Payment from "./Payment.svelte";

  let {
    product,
    close,
  }: {
    product: Product;
    close: () => void;
  } = $props();

  let loadingPayment: boolean = $state<boolean>(false);
  let justPurchased: boolean = $state(false);
  let popupUrl: string | null = $state<string | null>(null);
  let giftClass: string = $derived<string>(justPurchased ? "purchased" : "");

  async function handleClick(id: string) {
    loadingPayment = true;
    try {
      const { payment } = await api.getPaymentUrl(id);
      popupUrl = payment.url;
    } finally {
      loadingPayment = false;
    }
  }

  async function handlePurchase(purchase: Purchase | null) {
    justPurchased = !!purchase;
  }
</script>

{#if popupUrl}
  <Payment url={popupUrl} callback={handlePurchase} />
{/if}

<!-- svelte-ignore a11y_click_events_have_key_events, a11y_no_static_element_interactions -->
<div class="backdrop" transition:fade onclick={close}>
  <div class="gift" onclick={(e) => e.stopPropagation()}>
    <button class="close-btn" onclick={close}>&times;</button>

    <img src={product.imageUrl} alt={product.name} />
    <h2>{product.name}</h2>
    <p class="price">{formatPrice(product.priceBrl)}</p>

    {#if product.purchased}
      <p class="msg">This gift has already been purchased. Thank you!</p>
    {:else if !product.active}
      <p class="msg">This gift is not available.</p>
    {:else if justPurchased}
      <p class="msg purchased" transition:fade>Thanks for the purchase!</p>
    {:else}
      <button
        class="buy-btn"
        disabled={loadingPayment}
        onclick={() => handleClick(product.id)}
      >
        {loadingPayment ? "Generating Link..." : "Give this Gift"}
      </button>
    {/if}
  </div>
</div>

<style>
  .backdrop {
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

  .gift {
    background: white;
    width: 100%;
    max-width: 350px;
    border-radius: 16px;
    padding: 24px;
    position: relative;
    text-align: center;
    box-shadow: 0 10px 25px rgba(0, 0, 0, 0.2);

    transition: background 1s ease;
  }

  .purchased {
    border-radius: 1rem;
    padding: 0.5rem;
    background: greenyellow;
  }

  .gift img {
    width: 100%;
    aspect-ratio: 1;
    object-fit: cover;
    border-radius: 12px;
    margin-bottom: 16px;
  }

  .gift h2 {
    font-size: 1.25rem;
    margin: 0 0 8px 0;
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
