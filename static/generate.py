import numpy as np
import cv2

# Create a black image
img = np.zeros((400, 400 ,3), np.uint8)

# Write some Text
print(img.shape)
font                   = cv2.FONT_HERSHEY_SIMPLEX
fontScale              = 2
fontColor              = (255,255,255)
thickness              = 2
lineType               = 0

text = 'Immagine'

text_width, text_height = cv2.getTextSize(text, font, fontScale, thickness)[0]

CenterCoordinates = (int(img.shape[1] / 2) - int(text_width / 2), int(img.shape[0] / 2) + int(text_height / 2))

cv2.putText(img, text, CenterCoordinates, font, fontScale, fontColor, thickness)
cv2.imwrite('out.jpg', img)
# cv2.imshow('image', img)

# cv2.waitKey(0)
# cv2.destroyAllWindows()
# cv2.
