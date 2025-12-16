package com.floatisland.app;

import android.app.Notification;
import android.content.Context;
import android.content.Intent;
import android.service.notification.NotificationListenerService;
import android.service.notification.StatusBarNotification;
import android.util.Log;
import android.os.Bundle;
import android.widget.Toast;

import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class NotificationMonitorService extends NotificationListenerService {

    private static final String TAG = "Float_Monitor";
    private FloatingWindowHelper floatingWindowHelper;

    @Override
    public void onCreate() {
        super.onCreate();
        Log.i(TAG, ">>> Service Created <<<");
        try {
            floatingWindowHelper = new FloatingWindowHelper(getApplicationContext());
        } catch (Exception e) {
            Log.e(TAG, "Error creating helper", e);
        }
    }

    @Override
    public void onListenerConnected() {
        super.onListenerConnected();
        Log.i(TAG, ">>> Listener Connected <<<");
        showDebugToast("监听服务已连接");
    }

    @Override
    public void onListenerDisconnected() {
        super.onListenerDisconnected();
        Log.i(TAG, ">>> Listener Disconnected <<<");
        showDebugToast("监听服务断开");
    }

    @Override
    public void onNotificationPosted(StatusBarNotification sbn) {
        if (sbn == null)
            return;

        String packageName = sbn.getPackageName();
        // Filter out self
        if (getPackageName().equals(packageName))
            return;

        Notification notification = sbn.getNotification();
        if (notification == null || notification.extras == null) {
            return;
        }

        Bundle extras = notification.extras;
        CharSequence titleChar = extras.getCharSequence(Notification.EXTRA_TITLE);
        CharSequence textChar = extras.getCharSequence(Notification.EXTRA_TEXT);

        String title = titleChar != null ? titleChar.toString() : "";
        String text = textChar != null ? textChar.toString() : "";
        String combinedText = (title + " " + text).toLowerCase();

        // Check Mode
        boolean isTestMode = getSharedPreferences("FloatPrefs", Context.MODE_PRIVATE)
                .getBoolean("monitor_enabled", false);

        if (isTestMode) {
            // TEST MODE: Show Everything
            Log.i(TAG, "[Test Mode] Processing: " + packageName);
            showDebugToast("测试模式收到: " + packageName);
            safeShowWindow(text, title + "\n(" + packageName + ")");
        } else {
            // PRODUCTION MODE: Smart Filter
            if (isTransactionRelated(combinedText)) {
                Log.i(TAG, "[Smart Filter] Matched: " + packageName);
                safeShowWindow(text, title); // Clean display for production
            } else {
                Log.d(TAG, "[Smart Filter] Ignored non-transaction: " + packageName);
            }
        }
    }

    private boolean isTransactionRelated(String text) {
        if (text == null)
            return false;
        String[] keywords = {
                "支付", "收款", "到账", "消费", "交易", "余额", "转账", "付款",
                "payment", "transaction", "received", "spent", "支出"
        };
        for (String key : keywords) {
            if (text.contains(key))
                return true;
        }
        // Also check for currency symbols followed by numbers loosely
        return text.matches(".*(¥|\\$|€|元).*\\d+.*");
    }

    private String extractAmount(String text) {
        if (text == null)
            return "0.00";
        try {
            // Match number like 12.34
            Pattern p = Pattern.compile("\\d+(\\.\\d{2})?");
            Matcher m = p.matcher(text);
            if (m.find()) {
                return m.group();
            }
        } catch (Exception e) {
            Log.e(TAG, "Regex Failed", e);
        }
        return "0.00";
    }

    private void safeShowWindow(String amount, String merchant) {
        Log.i(TAG, "Preparing to show window...");
        try {
            if (floatingWindowHelper == null) {
                Log.w(TAG, "Helper was null, recreating...");
                floatingWindowHelper = new FloatingWindowHelper(getApplicationContext());
            }

            String detectedAmount = extractAmount(amount);

            // Logic to determine display text
            String displayAmount = "0.00";
            String displayMerchant = merchant;

            if (detectedAmount != null && !detectedAmount.equals("0.00")) {
                displayAmount = detectedAmount;
            } else {
                displayMerchant = merchant + "\n" + (amount != null ? amount : "");
            }

            Log.i(TAG, "Calling helper.showFloatingWindow(" + displayAmount + ")");
            floatingWindowHelper.showFloatingWindow(displayAmount, displayMerchant);

        } catch (Exception e) {
            Log.e(TAG, "CRITICAL ERROR in safeShowWindow", e);
            showDebugToast("ERROR: " + e.getMessage());
        }
    }

    private void showDebugToast(String msg) {
        try {
            new android.os.Handler(android.os.Looper.getMainLooper())
                    .post(() -> Toast.makeText(getApplicationContext(), "[Float] " + msg, Toast.LENGTH_SHORT).show());
        } catch (Exception e) {
            Log.e(TAG, "Toast failed", e);
        }
    }
}
