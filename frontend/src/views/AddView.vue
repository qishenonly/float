<script setup>
import { ref, onMounted, computed } from "vue";
import { useRouter } from "vue-router";
import { categoryAPI } from "@/api/category";
import { accountAPI } from "@/api/account";
import { transactionAPI } from "@/api/transaction";
import { useToast } from "../composables/useToast";

const router = useRouter();
const { showToast } = useToast();

const type = ref("expense"); // 'expense', 'income', or 'transfer'
const amount = ref("");
const description = ref("");
const date = ref(new Date().toISOString().split("T")[0]);
const loading = ref(false);
const showFromDropdown = ref(false);
const showToDropdown = ref(false);

// Auto-resize textarea
const textareaRef = ref(null);
const adjustTextareaHeight = () => {
  if (textareaRef.value) {
    textareaRef.value.style.height = "auto";
    const scrollHeight = textareaRef.value.scrollHeight;
    const maxHeight = textareaRef.value.scrollHeight; // Allow scrolling
    textareaRef.value.style.height = Math.min(scrollHeight, 96) + "px"; // 96px = ~4 lines
  }
};

// Categories
const expenseCategories = ref([]);
const incomeCategories = ref([]);
const transferCategories = ref([]);
const selectedCategory = ref(null);

// Accounts
const accounts = ref([]);
const selectedAccount = ref(null);

// Transfer specific
const transferFrom = ref(null);
const transferTo = ref(null);

onMounted(() => {
  loadData();
});

const loadData = async () => {
  loading.value = true;
  try {
    const [expenseRes, incomeRes, accountsRes] = await Promise.all([
      categoryAPI.getCategories("expense"),
      categoryAPI.getCategories("income"),
      accountAPI.getAccounts(),
    ]);

    expenseCategories.value = expenseRes.data || [];
    incomeCategories.value = incomeRes.data || [];
    accounts.value = accountsRes.data || [];

    // Set defaults
    if (expenseCategories.value.length > 0) {
      selectedCategory.value = expenseCategories.value[0];
    }
    if (accounts.value.length > 0) {
      selectedAccount.value = accounts.value[0];
      transferFrom.value = accounts.value[0];
      if (accounts.value.length > 1) {
        transferTo.value = accounts.value[1];
      }
    }


  } catch (error) {
    console.error("Failed to load data:", error);
    showToast("加载数据失败", "error");
  } finally {
    loading.value = false;
  }
};

const switchType = (newType) => {
  type.value = newType;
  if (newType === "income") {
    if (incomeCategories.value.length > 0) {
      selectedCategory.value = incomeCategories.value[0];
    }
  } else if (newType === "expense") {
    if (expenseCategories.value.length > 0) {
      selectedCategory.value = expenseCategories.value[0];
    }
  } else if (newType === "transfer") {
    if (transferCategories.value.length > 0) {
      selectedCategory.value = transferCategories.value[0];
    }
  }
};

const selectCategory = (cat) => {
  selectedCategory.value = cat;
};

const selectAccount = (acc) => {
  selectedAccount.value = acc;
};

const swapTransferAccounts = () => {
   [transferFrom.value, transferTo.value] = [
     transferTo.value,
     transferFrom.value,
   ];
};

const selectTransferFrom = (acc) => {
  transferFrom.value = acc;
  showFromDropdown.value = false;
};

const selectTransferTo = (acc) => {
  transferTo.value = acc;
  showToDropdown.value = false;
};

const getIconClasses = (account) => {
  const brandAccounts = ["alipay", "wechat"];
  const isBrand = brandAccounts.includes(account.account_type);
  return isBrand ? ["fa-brands", account.icon] : ["fa-solid", account.icon];
};

const buttonColor = computed(() => {
  switch (type.value) {
    case "income":
      return "bg-green-600 shadow-green-200";
    case "transfer":
      return "bg-blue-600 shadow-blue-200";
    default:
      return "bg-gray-800 shadow-gray-200";
  }
});

const indicatorColor = computed(() => {
  switch (type.value) {
    case "income":
      return "bg-green-500";
    case "transfer":
      return "bg-blue-500";
    default:
      return "bg-orange-400";
  }
});

const tabTextColor = computed(() => {
  switch (type.value) {
    case "income":
      return "text-green-600";
    case "transfer":
      return "text-blue-600";
    default:
      return "text-gray-800";
  }
});

const currentCategories = computed(() => {
  if (type.value === "expense") return expenseCategories.value;
  if (type.value === "income") return incomeCategories.value;
  return transferCategories.value;
});

const handleSave = async () => {
  // 验证
  if (!amount.value || parseFloat(amount.value) <= 0) {
    showToast("请输入有效的金额", "warning");
    return;
  }

  if (!selectedCategory.value?.id) {
    showToast("请选择分类", "warning");
    return;
  }

  if (!selectedAccount.value?.id && type.value !== "transfer") {
    showToast("请选择账户", "warning");
    return;
  }

  if (type.value === "transfer") {
    if (!transferFrom.value?.id || !transferTo.value?.id) {
      showToast("请选择转账账户", "warning");
      return;
    }
  }

  loading.value = true;
  try {
    const payload = {
       type: type.value,
       category_id: selectedCategory.value.id,
       amount: parseFloat(amount.value),
       currency: "CNY",
       title: selectedCategory.value.name,
       description: description.value,
       transaction_date: `${date.value}T00:00:00Z`,
     };

    // 根据交易类型填充不同字段
    if (type.value === "transfer") {
      payload.account_id = transferFrom.value.id;
      payload.to_account_id = transferTo.value.id;
    } else {
      payload.account_id = selectedAccount.value.id;
    }

    await transactionAPI.createTransaction(payload);
    showToast("交易记录保存成功", "success");

    // 延迟后返回首页
    setTimeout(() => {
      router.push("/");
    }, 500);
  } catch (error) {
    console.error("Failed to save transaction:", error);
    showToast(error.response?.data?.message || "保存失败", "error");
  } finally {
    loading.value = false;
  }
};
</script>

<template>
  <div class="w-full flex flex-col">
    <!-- Header -->
    <div
      class="px-6 pt-10 pb-4 flex justify-between items-center flex-shrink-0"
    >
      <RouterLink
        to="/"
        class="w-10 h-10 flex items-center justify-center text-gray-400 hover:bg-white rounded-full transition shadow-sm bg-white/50 backdrop-blur-md"
      >
        <i class="fa-solid fa-xmark text-lg"></i>
      </RouterLink>

      <div class="flex bg-gray-100/80 p-1.5 rounded-2xl backdrop-blur-md">
        <button
          @click="switchType('expense')"
          class="tab-btn px-5 py-2 rounded-xl text-xs font-bold transition-all duration-300"
          :class="
            type === 'expense'
              ? 'bg-white shadow-sm text-gray-800 transform scale-105'
              : 'text-gray-600 font-medium hover:bg-white/50'
          "
        >
          支出
        </button>
        <button
          @click="switchType('income')"
          class="tab-btn px-5 py-2 rounded-xl text-xs font-bold transition-all duration-300"
          :class="
            type === 'income'
              ? 'bg-white shadow-sm text-green-600 transform scale-105'
              : 'text-gray-600 font-medium hover:bg-white/50'
          "
        >
          收入
        </button>
        <button
          @click="switchType('transfer')"
          class="tab-btn px-5 py-2 rounded-xl text-xs font-bold transition-all duration-300"
          :class="
            type === 'transfer'
              ? 'bg-white shadow-sm text-blue-600 transform scale-105'
              : 'text-gray-600 font-medium hover:bg-white/50'
          "
        >
          转账
        </button>
      </div>

      <div class="w-10"></div>
    </div>

    <!-- Amount Input Section -->
    <div class="px-6 py-6 flex flex-col items-end relative z-10 flex-shrink-0">
      <div
        v-if="selectedCategory || selectedAccount"
        class="flex items-center gap-2 mb-1"
      >
        <div
          class="bg-white/60 px-3 py-1 rounded-full flex items-center gap-2 shadow-sm border border-white/50 backdrop-blur-sm"
        >
          <div class="w-2 h-2 rounded-full" :class="indicatorColor"></div>
          <span class="text-[10px] font-bold text-gray-600">
            {{
              type === "transfer" ? "转账" : selectedCategory?.name || "未分类"
            }}
            · {{ selectedAccount?.account_name || "未选账户" }}
          </span>
        </div>
      </div>
      <div
        class="flex items-baseline w-full justify-end gap-2 text-gray-800 relative"
      >
        <span class="text-3xl font-bold -mb-2 opacity-60">¥</span>
        <input
          v-model="amount"
          type="number"
          placeholder="0.00"
          class="w-full text-right text-[3.5rem] font-extrabold bg-transparent outline-none placeholder-gray-200 caret-blue-500 leading-tight tracking-tight"
          :class="type === 'income' ? 'text-green-600' : 'text-gray-800'"
          autofocus
        />
      </div>
    </div>

    <!-- Content Area -->
    <div class="overflow-hidden px-6">
      <!-- Standard Section (Expense/Income) -->
      <div
        v-if="type !== 'transfer'"
        class="py-6 transition-opacity duration-300"
      >
        <!-- Account Selection -->
        <div class="mb-8">
          <p
            class="text-[10px] text-gray-400 font-bold mb-4 px-1 uppercase tracking-widest"
          >
            选择账户
          </p>
          <div class="flex flex-wrap gap-3">
            <div
              v-for="acc in accounts"
              :key="acc.id"
              @click="selectAccount(acc)"
              class="px-4 py-3 rounded-full text-xs font-bold flex items-center gap-2.5 cursor-pointer shadow-sm transition-all duration-300 border"
              :class="
                selectedAccount?.id === acc.id
                  ? 'bg-gray-900 text-white border-gray-900 shadow-md'
                  : 'bg-white border-gray-200 text-gray-600 hover:border-gray-300'
              "
            >
              <div
                class="w-6 h-6 rounded-full flex items-center justify-center text-sm flex-shrink-0"
                :class="
                  selectedAccount?.id === acc.id
                    ? 'bg-white/20'
                    : `bg-${acc.color}-100 text-${acc.color}-600`
                "
              >
                <i :class="getIconClasses(acc)"></i>
              </div>
              {{ acc.account_name }}
            </div>
          </div>
        </div>

        <!-- Category Selection -->
        <div class="mb-8">
          <p
            class="text-[10px] text-gray-400 font-bold mb-4 px-1 uppercase tracking-widest"
          >
            选择分类
          </p>
          <div class="grid grid-cols-5 gap-y-6 gap-x-2">
            <div
              v-for="cat in currentCategories"
              :key="cat.id"
              @click="selectCategory(cat)"
              class="category-item flex flex-col items-center gap-2 cursor-pointer transition-all duration-300"
              :class="selectedCategory?.id === cat.id ? 'active' : ''"
            >
              <div
                class="icon-box w-[3.25rem] h-[3.25rem] rounded-[1.2rem] flex items-center justify-center text-xl shadow-sm border transition-all duration-300"
                :class="
                  selectedCategory?.id === cat.id
                    ? `bg-gray-900 text-white shadow-lg transform -translate-y-1 scale-110`
                    : `bg-${cat.color}-50 text-${cat.color}-500 border-${cat.color}-100 border-opacity-50`
                "
              >
                <i class="fa-solid" :class="cat.icon"></i>
              </div>
              <span
                class="text-[10px] font-medium transition-colors duration-200"
                :class="
                  selectedCategory?.id === cat.id
                    ? 'text-gray-900 font-bold'
                    : 'text-gray-400'
                "
              >
                {{ cat.name }}
              </span>
            </div>
          </div>
        </div>

        <!-- Notes Section -->
        <div class="glass-input rounded-2xl p-4 flex flex-col gap-3">
          <div class="flex items-start gap-3">
            <div
              class="mt-1 w-6 h-6 rounded-full bg-gray-100 flex items-center justify-center text-gray-400"
            >
              <i class="fa-solid fa-pen text-xs"></i>
            </div>
            <textarea
              ref="textareaRef"
              v-model="description"
              @input="adjustTextareaHeight"
              placeholder="添加备注..."
              class="w-full text-sm outline-none text-gray-700 bg-transparent resize-none placeholder-gray-400 font-medium overflow-y-auto"
              style="min-height: 24px; max-height: 96px"
            ></textarea>
          </div>
          <div class="h-px bg-gray-100 w-full my-1"></div>
          <div class="flex justify-between items-center">
            <div class="flex gap-4">
              <div
                class="flex items-center gap-1.5 text-gray-500 hover:text-gray-800 cursor-pointer transition group"
              >
                <i
                  class="fa-regular fa-calendar-check text-sm group-hover:text-blue-500"
                ></i>
                <span class="text-[10px] font-bold">今天</span>
              </div>
              <div
                class="flex items-center gap-1.5 text-gray-500 hover:text-gray-800 cursor-pointer transition group"
              >
                <i
                  class="fa-solid fa-hashtag text-sm group-hover:text-blue-500"
                ></i>
                <span class="text-[10px] font-bold">标签</span>
              </div>
              <div
                class="flex items-center gap-1.5 text-gray-500 hover:text-gray-800 cursor-pointer transition group"
              >
                <i
                  class="fa-regular fa-image text-sm group-hover:text-blue-500"
                ></i>
              </div>
              <div
                class="flex items-center gap-1.5 text-gray-500 hover:text-gray-800 cursor-pointer transition group"
              >
                <i
                  class="fa-solid fa-location-dot text-sm group-hover:text-blue-500"
                ></i>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Transfer Section -->
      <div v-else class="py-6 transition-opacity duration-300">
        <div class="mb-8">
          <!-- Transfer From Account Selection -->
          <div class="mb-1">
            <p class="text-[10px] text-gray-400 font-bold mb-2 px-1 uppercase tracking-wider">转出 (FROM)</p>
            <div class="relative">
              <button
                @click="showFromDropdown = !showFromDropdown"
                class="w-full px-4 py-3 rounded-t-2xl bg-white border border-gray-200 border-b-0 flex items-center justify-between cursor-pointer shadow-sm hover:border-gray-300 transition-all duration-300"
              >
                <div class="flex items-center gap-3">
                  <div
                    class="w-6 h-6 rounded-full flex items-center justify-center text-sm flex-shrink-0"
                    :class="`bg-${transferFrom?.color}-100 text-${transferFrom?.color}-600`"
                  >
                    <i :class="getIconClasses(transferFrom)"></i>
                  </div>
                  <span class="text-sm font-bold text-gray-800">{{ transferFrom?.account_name }}</span>
                </div>
                <i class="fa-solid fa-chevron-right text-gray-400 text-xs"></i>
              </button>
              
              <!-- Dropdown Menu -->
              <div
                v-if="showFromDropdown"
                class="absolute top-full left-0 right-0 bg-white border border-gray-200 rounded-b-2xl shadow-lg z-50 max-h-60 overflow-y-auto"
              >
                <div
                  v-for="acc in accounts"
                  :key="`from-${acc.id}`"
                  @click="selectTransferFrom(acc)"
                  class="px-4 py-2.5 cursor-pointer hover:bg-gray-50 flex items-center gap-3 border-b border-gray-100 last:border-b-0"
                >
                  <div
                    class="w-5 h-5 rounded-full flex items-center justify-center text-xs flex-shrink-0"
                    :class="`bg-${acc.color}-100 text-${acc.color}-600`"
                  >
                    <i :class="getIconClasses(acc)"></i>
                  </div>
                  <span class="text-xs font-bold text-gray-800">{{ acc.account_name }}</span>
                </div>
              </div>
            </div>
          </div>

          <!-- Swap Button in Middle -->
          <div class="flex justify-center -my-5 relative z-10">
            <button
              @click="swapTransferAccounts"
              class="w-10 h-10 rounded-full bg-white text-gray-600 shadow-md flex items-center justify-center border border-gray-200 hover:rotate-180 transition duration-300"
            >
              <i class="fa-solid fa-arrow-right-arrow-left fa-rotate-90 text-xs"></i>
            </button>
          </div>

          <!-- Transfer To Account Selection -->
          <div class="mt-1">
            <p class="text-[10px] text-gray-400 font-bold mb-2 px-1 uppercase tracking-wider">转入 (TO)</p>
            <div class="relative">
              <button
                @click="showToDropdown = !showToDropdown"
                class="w-full px-4 py-3 rounded-b-2xl bg-white border border-gray-200 border-t-0 flex items-center justify-between cursor-pointer shadow-sm hover:border-gray-300 transition-all duration-300"
              >
                <div class="flex items-center gap-3">
                  <div
                    class="w-6 h-6 rounded-full flex items-center justify-center text-sm flex-shrink-0"
                    :class="`bg-${transferTo?.color}-100 text-${transferTo?.color}-600`"
                  >
                    <i :class="getIconClasses(transferTo)"></i>
                  </div>
                  <span class="text-sm font-bold text-gray-800">{{ transferTo?.account_name }}</span>
                </div>
                <i class="fa-solid fa-chevron-right text-gray-400 text-xs"></i>
              </button>
              
              <!-- Dropdown Menu -->
              <div
                v-if="showToDropdown"
                class="absolute top-full left-0 right-0 bg-white border border-gray-200 rounded-b-2xl shadow-lg z-50 max-h-60 overflow-y-auto"
              >
                <div
                  v-for="acc in accounts"
                  :key="`to-${acc.id}`"
                  @click="selectTransferTo(acc)"
                  class="px-4 py-2.5 cursor-pointer hover:bg-gray-50 flex items-center gap-3 border-b border-gray-100 last:border-b-0"
                >
                  <div
                    class="w-5 h-5 rounded-full flex items-center justify-center text-xs flex-shrink-0"
                    :class="`bg-${acc.color}-100 text-${acc.color}-600`"
                  >
                    <i :class="getIconClasses(acc)"></i>
                  </div>
                  <span class="text-xs font-bold text-gray-800">{{ acc.account_name }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Transfer Notes -->
        <div class="glass-input rounded-2xl p-4 flex flex-col gap-3 mt-6">
          <div class="flex items-start gap-3">
            <div
              class="mt-1 w-6 h-6 rounded-full bg-gray-100 flex items-center justify-center text-gray-400"
            >
              <i class="fa-solid fa-pen text-xs"></i>
            </div>
            <textarea
              ref="textareaRef"
              v-model="description"
              @input="adjustTextareaHeight"
              placeholder="添加备注..."
              class="w-full text-sm outline-none text-gray-700 bg-transparent resize-none placeholder-gray-400 font-medium overflow-y-auto"
              style="min-height: 24px; max-height: 96px"
            ></textarea>
          </div>
          <div class="h-px bg-gray-100 w-full my-1"></div>
          <div class="flex justify-between items-center">
            <div class="flex gap-4">
              <div
                class="flex items-center gap-1.5 text-gray-500 hover:text-gray-800 cursor-pointer transition group"
              >
                <i
                  class="fa-regular fa-calendar-check text-sm group-hover:text-blue-500"
                ></i>
                <span class="text-[10px] font-bold">今天</span>
              </div>
              <div
                class="flex items-center gap-1.5 text-gray-500 hover:text-gray-800 cursor-pointer transition group"
              >
                <i
                  class="fa-solid fa-hashtag text-sm group-hover:text-blue-500"
                ></i>
                <span class="text-[10px] font-bold">标签</span>
              </div>
              <div
                class="flex items-center gap-1.5 text-gray-500 hover:text-gray-800 cursor-pointer transition group"
              >
                <i
                  class="fa-regular fa-image text-sm group-hover:text-blue-500"
                ></i>
              </div>
              <div
                class="flex items-center gap-1.5 text-gray-500 hover:text-gray-800 cursor-pointer transition group"
              >
                <i
                  class="fa-solid fa-location-dot text-sm group-hover:text-blue-500"
                ></i>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Action Bar - Fixed at bottom -->
    <div
      class="flex-shrink-0 px-6 py-4"
    >
      <button
        @click="handleSave"
        :disabled="loading"
        class="w-full text-white font-bold text-lg h-12 rounded-full active:scale-[0.98] transition flex items-center justify-center gap-3 group disabled:opacity-60"
        :class="buttonColor"
      >
        <span>{{ loading ? "保存中..." : "保存一笔" }}</span>
        <i
          v-if="!loading"
          class="fa-solid fa-arrow-right text-sm opacity-50 group-hover:translate-x-1 transition"
        ></i>
        <i v-else class="fa-solid fa-spinner fa-spin text-sm opacity-50"></i>
      </button>
    </div>
  </div>
</template>

<style scoped>
:root {
  --color-blue: #3b82f6;
  --color-red: #ef4444;
  --color-green: #10b981;
  --color-purple: #a855f7;
  --color-orange: #f97316;
  --color-pink: #ec4899;
  --color-yellow: #eab308;
}

.hide-scrollbar::-webkit-scrollbar {
  display: none;
}

/* Glass input effect */
.glass-input {
  background: rgba(255, 255, 255, 0.7);
  backdrop-filter: blur(12px);
  border: 1px solid rgba(255, 255, 255, 0.8);
  box-shadow: 0 4px 10px -2px rgba(0, 0, 0, 0.03);
  transition: all 0.2s ease-in-out;
}

.glass-input:active {
  transform: scale(0.99);
}

/* Tab button animation */
.tab-btn {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

/* Category item animation */
.category-item .icon-box {
  transition: all 0.3s cubic-bezier(0.34, 1.56, 0.64, 1);
}

/* Account capsule */
.account-capsule {
  transition: all 0.2s ease-in-out;
}

.account-capsule.active {
  background-color: #1f2937 !important;
  color: #ffffff !important;
  box-shadow: 0 4px 10px rgba(31, 41, 55, 0.3);
  transform: translateY(-1px);
}

/* Hide number input spinners */
input[type="number"]::-webkit-inner-spin-button,
input[type="number"]::-webkit-outer-spin-button {
  -webkit-appearance: none;
  margin: 0;
}
</style>
