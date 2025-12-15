package com.floatisland.app;

import android.app.Notification;
import android.service.notification.NotificationListenerService;
import android.service.notification.StatusBarNotification;
import android.util.Log;
import android.os.Bundle;

import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class NotificationMonitorService extends NotificationListenerService {

    private static final String TAG = "NotificationMonitor";
    private FloatingWindowHelper floatingWindowHelper;

    @Override
    public void onCreate() {
        super.onCreate();
        floatingWindowHelper = new FloatingWindowHelper(this);
    }

    @Override
    public void onNotificationPosted(StatusBarNotification sbn) {
        // FOR TESTING: Listen to ALL notifications
        String packageName = sbn.getPackageName();
        Notification notification = sbn.getNotification();
        if (notification == null)
            return;

        Bundle extras = notification.extras;
        if (extras == null)
            return;

        String title = extras.getString(Notification.EXTRA_TITLE, "");
        String text = extras.getString(Notification.EXTRA_TEXT, "");

        Log.d(TAG, "Notification received from " + packageName + ": " + title + " : " + text);

        // Filter out system or irrelevant notifications
        // We skip our own app to avoid loops if we post notifications
        if (packageName.equals(getPackageName()))
            return;

        // Filtering logic:
        boolean valid = false;
        if (packageName.equals("com.tencent.mm") || packageName.equals("com.eg.android.AlipayGphone")) {
            if (title.contains("支付") || title.contains("收款") || title.contains("到账") ||
                    text.contains("支付") || text.contains("收款") || text.contains("到账") ||
                    text.contains("元") || text.contains("￥") || text.contains("¥")) {
                valid = true;
            }
        }

        // For other apps (banks etc), we might want broader checks or explicit package
        // additions
        // But user asked "if notification is payment/receipt message"

        if (!valid) {
            // Check for generic payment keywords even if seemingly not from main apps,
            // but to reduce spam, we might restrict it.
            // Let's rely on money symbols or keywords for now for ALL apps if not
            // specifically excluded
            if (title.contains("支付") || title.contains("收款") || title.contains("到账") ||
                    text.contains("支付") || text.contains("收款") || text.contains("到账") ||
                    (text.contains("元") && text.matches(".*\\d+.*"))) {
                valid = true;
            }
        }

        if (!valid)
            return;

        // Show window
        showWindow(text, title + " (" + packageName + ")");
    }

    private String extractAmount(String text) {
        if (text == null)
            return null;
        // Regex for amount
        Pattern p = Pattern.compile("\\d+(\\.\\d{2})?");
        Matcher m = p.matcher(text);
        if (m.find()) {
            return m.group();
        }
        return "0.00"; // For testing, ensure we have a value
    }

    private void showWindow(String amount, String merchant) {
        // For testing, amount might be the text content
        // If amount extraction fails, we just show the text to verify we caught the
        // notification
        String detectedAmount = extractAmount(amount);
        if (detectedAmount == null || detectedAmount.equals("0.00")) {
            // If no amount found, just put 0 but show text in description
            floatingWindowHelper.showFloatingWindow("0.00", merchant + "\n" + amount);
        } else {
            floatingWindowHelper.showFloatingWindow(detectedAmount, merchant);
        }
    }
}
