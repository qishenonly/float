package com.floatisland.app;

import android.accessibilityservice.AccessibilityService;
import android.view.accessibility.AccessibilityEvent;
import android.view.accessibility.AccessibilityNodeInfo;
import android.util.Log;
import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class AutoBookkeepingService extends AccessibilityService {

    private static final String TAG = "AutoBookkeepingService";
    private FloatingWindowHelper floatingWindowHelper;
    private long lastProcessedTime = 0;
    private static final long COOLDOWN_MS = 3000; // Prevent duplicate popups

    @Override
    public void onCreate() {
        super.onCreate();
        floatingWindowHelper = new FloatingWindowHelper(this);
    }

    @Override
    public void onAccessibilityEvent(AccessibilityEvent event) {
        if (event == null)
            return;

        // Debounce
        long currentTime = System.currentTimeMillis();
        if (currentTime - lastProcessedTime < COOLDOWN_MS) {
            return;
        }

        String packageName = event.getPackageName() != null ? event.getPackageName().toString() : "";

        // Ensure we are on the correct package
        if (!packageName.equals("com.eg.android.AlipayGphone") && !packageName.equals("com.tencent.mm")) {
            return;
        }

        AccessibilityNodeInfo rootNode = getRootInActiveWindow();
        if (rootNode == null)
            return;

        // Process Alipay
        if (packageName.equals("com.eg.android.AlipayGphone")) {
            processAlipay(rootNode);
        }
        // Process WeChat
        else if (packageName.equals("com.tencent.mm")) {
            processWeChat(rootNode);
        }
    }

    private void processAlipay(AccessibilityNodeInfo rootNode) {
        // Look for "支付成功", "交易成功", "付款成功"
        if (matchesText(rootNode, "支付成功|交易成功|付款成功")) {
            // Try to find amount.
            // Alipay typically shows amount like "-12.50" or "12.50" in a large font
            String amount = findAmount(rootNode);
            String merchant = findMerchant(rootNode);

            if (amount != null) {
                showWindow(amount, merchant);
            }
        }
    }

    private void processWeChat(AccessibilityNodeInfo rootNode) {
        // Look for "支付成功", "完成", "付款成功"
        if (matchesText(rootNode, "支付成功|完成|付款成功")) {
            String amount = findAmount(rootNode);
            String merchant = findMerchant(rootNode);

            if (amount != null) {
                showWindow(amount, merchant);
            }
        }
    }

    private boolean matchesText(AccessibilityNodeInfo node, String pattern) {
        // pattern unused in this simple implementation, we check specifically
        if (!node.findAccessibilityNodeInfosByText("支付成功").isEmpty())
            return true;
        if (!node.findAccessibilityNodeInfosByText("交易成功").isEmpty())
            return true;
        if (!node.findAccessibilityNodeInfosByText("付款成功").isEmpty())
            return true;
        if (!node.findAccessibilityNodeInfosByText("完成").isEmpty())
            return true;
        return false;
    }

    // Recursive search or keyword search for amount
    private String findAmount(AccessibilityNodeInfo node) {
        // Simplified strategy: look for text matching money format
        // This is a naive implementation and should be refined with specific view IDs
        // if possible,
        // but view IDs change with app updates. Regex is more robust.

        // Regex for currency:
        // ¥ 12.50
        // -12.50
        // 12.50
        Pattern p = Pattern.compile("[-+]?\\d{1,3}(?:,\\d{3})*\\.\\d{2}");

        // Perform a breadth-first or depth-first search
        // For efficiency, we might just look for nodes that contain text like digits

        return searchAmountRecursive(node, p);
    }

    private String searchAmountRecursive(AccessibilityNodeInfo node, Pattern p) {
        if (node == null)
            return null;

        if (node.getText() != null) {
            String text = node.getText().toString();
            Matcher m = p.matcher(text);
            if (m.find()) {
                return m.group();
            }
        }

        for (int i = 0; i < node.getChildCount(); i++) {
            String result = searchAmountRecursive(node.getChild(i), p);
            if (result != null)
                return result;
        }
        return null;
    }

    private String findMerchant(AccessibilityNodeInfo node) {
        // Harder to generalize. Users often edit this anyway.
        // For now, return "扫码支付" or similar default if not found
        // Optimized: try to find text near "收款方" or similar labels
        return "自动识别交易";
    }

    private void showWindow(String amount, String merchant) {
        Log.i(TAG, "Detected payment: " + amount);
        lastProcessedTime = System.currentTimeMillis();
        // Run on UI thread? AccessibilityService runs on main thread usually.
        floatingWindowHelper.showFloatingWindow(amount, merchant);
    }

    @Override
    public void onInterrupt() {
        Log.e(TAG, "Service Interrupted");
    }

    @Override
    public void onDestroy() {
        super.onDestroy();
        if (floatingWindowHelper != null) {
            floatingWindowHelper.removeFloatingWindow();
        }
    }
}
