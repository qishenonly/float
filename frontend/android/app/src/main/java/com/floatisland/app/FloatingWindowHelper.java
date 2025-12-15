package com.floatisland.app;

import android.content.Context;
import android.content.Intent;
import android.graphics.PixelFormat;
import android.net.Uri;
import android.os.Build;
import android.view.Gravity;
import android.view.LayoutInflater;
import android.view.View;
import android.view.WindowManager;
import android.widget.Button;
import android.widget.EditText;
import android.widget.ImageView;
import android.widget.TextView;
import android.widget.Toast;

import java.util.Locale;

public class FloatingWindowHelper {

    private final Context context;
    private final WindowManager windowManager;
    private View floatingView;

    public FloatingWindowHelper(Context context) {
        this.context = context;
        this.windowManager = (WindowManager) context.getSystemService(Context.WINDOW_SERVICE);
    }

    public void showFloatingWindow(String amount, String merchant) {
        if (floatingView != null) {
            // Already showing, just update
            updateView(amount, merchant);
            return;
        }

        LayoutInflater inflater = LayoutInflater.from(context);
        floatingView = inflater.inflate(R.layout.floating_window, null);

        final EditText etAmount = floatingView.findViewById(R.id.et_amount);
        final EditText etDescription = floatingView.findViewById(R.id.et_description);
        Button btnSave = floatingView.findViewById(R.id.btn_save);
        // Button btnCancel = floatingView.findViewById(R.id.btn_cancel); // Removed in
        // new layout
        TextView tvTitle = floatingView.findViewById(R.id.tv_title);
        ImageView btnClose = floatingView.findViewById(R.id.btn_close_icon);

        if (btnClose != null) {
            btnClose.setOnClickListener(v -> removeFloatingWindow());
        }

        if (amount != null) {
            etAmount.setText(amount);
        }
        if (merchant != null) {
            etDescription.setText(merchant);
        }

        // Window params
        int layoutFlag;
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
            layoutFlag = WindowManager.LayoutParams.TYPE_APPLICATION_OVERLAY;
        } else {
            layoutFlag = WindowManager.LayoutParams.TYPE_PHONE;
        }

        final WindowManager.LayoutParams params = new WindowManager.LayoutParams(
                WindowManager.LayoutParams.WRAP_CONTENT,
                WindowManager.LayoutParams.WRAP_CONTENT,
                layoutFlag,
                WindowManager.LayoutParams.FLAG_NOT_FOCUSABLE | WindowManager.LayoutParams.FLAG_WATCH_OUTSIDE_TOUCH,
                PixelFormat.TRANSLUCENT);

        // Allow input (focusable)
        params.flags = WindowManager.LayoutParams.FLAG_NOT_TOUCH_MODAL
                | WindowManager.LayoutParams.FLAG_WATCH_OUTSIDE_TOUCH;

        params.gravity = Gravity.CENTER;

        try {
            windowManager.addView(floatingView, params);
        } catch (Exception e) {
            e.printStackTrace();
            Toast.makeText(context, "无法显示悬浮窗，请检查权限", Toast.LENGTH_SHORT).show();
            return;
        }

        // btnCancel.setOnClickListener(v -> removeFloatingWindow()); // Removed

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

        if (amount != null && !amount.isEmpty()) {
            etAmount.setText(amount);
        }
        if (merchant != null && !merchant.isEmpty()) {
            etDescription.setText(merchant);
        }
    }

    public void removeFloatingWindow() {
        if (floatingView != null) {
            try {
                windowManager.removeView(floatingView);
            } catch (Exception e) {
                e.printStackTrace();
            }
            floatingView = null;
        }
    }

    private void saveAndOpenApp(String amount, String description) {
        // Build Deep Link URI
        // Using "float://add" schema as defined in AndroidManifest
        // Format: float://add?amount=12.50&description=Starbucks

        Uri.Builder builder = new Uri.Builder();
        builder.scheme("float")
                .authority("add")
                .appendQueryParameter("amount", amount)
                .appendQueryParameter("description", description)
                // Add timestamp or other fields if needed
                .appendQueryParameter("auto", "true");

        Intent intent = new Intent(Intent.ACTION_VIEW, builder.build());
        intent.addFlags(Intent.FLAG_ACTIVITY_NEW_TASK);

        try {
            context.startActivity(intent);
        } catch (Exception e) {
            e.printStackTrace();
            Toast.makeText(context, "无法启动应用", Toast.LENGTH_SHORT).show();
        }
    }
}
