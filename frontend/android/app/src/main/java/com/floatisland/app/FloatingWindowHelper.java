package com.floatisland.app;

import android.content.Context;
import android.content.Intent;
import android.graphics.PixelFormat;
import android.net.Uri;
import android.os.Build;
import android.provider.Settings;
import android.util.Log;
import android.view.Gravity;
import android.view.LayoutInflater;
import android.view.View;
import android.view.WindowManager;
import android.widget.Button;
import android.widget.EditText;
import android.widget.ImageView;
import android.widget.TextView;
import android.widget.Toast;

public class FloatingWindowHelper {

    private static final String TAG = "Float_WindowHelper";
    private final Context context;
    private final WindowManager windowManager;
    private View floatingView;

    public FloatingWindowHelper(Context context) {
        // Use application context to avoid leaks and activity lifecycle issues
        this.context = context.getApplicationContext();
        this.windowManager = (WindowManager) this.context.getSystemService(Context.WINDOW_SERVICE);
    }

    public void showFloatingWindow(String amount, String merchant) {
        Log.i(TAG, "showFloatingWindow Called. Amount: " + amount);

        if (!checkOverlayPermission()) {
            return;
        }

        try {
            if (floatingView != null) {
                Log.i(TAG, "View already exists, updating...");
                updateView(amount, merchant);
                return;
            }

            Log.i(TAG, "Inflating view...");
            LayoutInflater inflater = LayoutInflater.from(context);
            // Pass null as root is correct for WindowManager
            floatingView = inflater.inflate(R.layout.floating_window, null);

            if (floatingView == null) {
                Log.e(TAG, "View Inflation Failed (result is null)");
                showToast("View Inflation Failed!");
                return;
            }

            setupViewElements(amount, merchant);

            WindowManager.LayoutParams params = createLayoutParams();

            Log.i(TAG, "Adding view to WindowManager...");
            windowManager.addView(floatingView, params);
            Log.i(TAG, "View added successfully.");
            showToast("悬浮窗已开启");

        } catch (Exception e) {
            Log.e(TAG, "Error showing floating window", e);
            showToast("显示失败: " + e.getClass().getSimpleName() + " - " + e.getMessage());
            e.printStackTrace();
            if (floatingView != null) {
                try {
                    windowManager.removeView(floatingView);
                } catch (Exception ex) {
                }
                floatingView = null;
            }
        }
    }

    private boolean checkOverlayPermission() {
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.M) {
            if (!Settings.canDrawOverlays(context)) {
                Log.w(TAG, "Overlay Permission MISSING");
                showToast("请开启[显示在其他应用上层]权限");
                requestOverlayPermission();
                return false;
            }
        }
        return true;
    }

    private void requestOverlayPermission() {
        try {
            Intent intent = new Intent(Settings.ACTION_MANAGE_OVERLAY_PERMISSION,
                    Uri.parse("package:" + context.getPackageName()));
            intent.setFlags(Intent.FLAG_ACTIVITY_NEW_TASK);
            context.startActivity(intent);
        } catch (Exception e) {
            Log.e(TAG, "Failed to open permission settings", e);
        }
    }

    private WindowManager.LayoutParams createLayoutParams() {
        int layoutFlag;
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
            layoutFlag = WindowManager.LayoutParams.TYPE_APPLICATION_OVERLAY;
        } else {
            layoutFlag = WindowManager.LayoutParams.TYPE_PHONE;
        }

        WindowManager.LayoutParams params = new WindowManager.LayoutParams(
                WindowManager.LayoutParams.WRAP_CONTENT,
                WindowManager.LayoutParams.WRAP_CONTENT,
                layoutFlag,
                // STATUS_BAR_PANEL allows input while keeping it as a floating window?
                // Actually usage of FLAG_NOT_TOUCH_MODAL is best for small windows that accept
                // input.
                // FLAG_WATCH_OUTSIDE_TOUCH allows dismissing when clicking outside (if
                // implemented)
                WindowManager.LayoutParams.FLAG_NOT_TOUCH_MODAL | WindowManager.LayoutParams.FLAG_WATCH_OUTSIDE_TOUCH,
                PixelFormat.TRANSLUCENT);

        params.gravity = Gravity.CENTER;
        return params;
    }

    private void setupViewElements(String amount, String merchant) {
        final EditText etAmount = floatingView.findViewById(R.id.et_amount);
        final EditText etDescription = floatingView.findViewById(R.id.et_description);
        Button btnSave = floatingView.findViewById(R.id.btn_save);
        ImageView btnClose = floatingView.findViewById(R.id.btn_close_icon);

        if (btnClose != null) {
            btnClose.setOnClickListener(v -> removeFloatingWindow());
        }

        if (amount != null)
            etAmount.setText(amount);
        if (merchant != null)
            etDescription.setText(merchant);

        btnSave.setOnClickListener(v -> {
            String savedAmount = etAmount.getText().toString();
            String savedDesc = etDescription.getText().toString();
            saveAndOpenApp(savedAmount, savedDesc);
            removeFloatingWindow();
        });
    }

    private void updateView(String amount, String merchant) {
        if (floatingView == null)
            return;
        EditText etAmount = floatingView.findViewById(R.id.et_amount);
        EditText etDescription = floatingView.findViewById(R.id.et_description);

        if (amount != null && !amount.isEmpty())
            etAmount.setText(amount);
        if (merchant != null && !merchant.isEmpty())
            etDescription.setText(merchant);

        showToast("内容已更新");
    }

    public void removeFloatingWindow() {
        if (floatingView != null) {
            try {
                windowManager.removeView(floatingView);
                Log.i(TAG, "Window removed");
            } catch (Exception e) {
                Log.e(TAG, "Error removing window", e);
            }
            floatingView = null;
        }
    }

    private void saveAndOpenApp(String amount, String description) {
        Uri.Builder builder = new Uri.Builder();
        builder.scheme("float")
                .authority("add")
                .appendQueryParameter("amount", amount)
                .appendQueryParameter("description", description)
                .appendQueryParameter("auto", "true");

        Intent intent = new Intent(Intent.ACTION_VIEW, builder.build());
        intent.addFlags(Intent.FLAG_ACTIVITY_NEW_TASK);

        try {
            context.startActivity(intent);
        } catch (Exception e) {
            Log.e(TAG, "Failed to launch app", e);
            showToast("无法启动应用");
        }
    }

    private void showToast(String msg) {
        try {
            new android.os.Handler(android.os.Looper.getMainLooper())
                    .post(() -> Toast.makeText(context, msg, Toast.LENGTH_SHORT).show());
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
